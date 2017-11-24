package beanstalkd

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
)

type Conn struct {
	conn net.Conn
	br   *bufio.Reader
	bw   *bufio.Writer
}

func Dial(network, dns string) *Conn {
	conn, err := net.Dial(network, dns)
	if err != nil {
		panic(err)
	}

	c := new(Conn)
	c.conn = conn
	c.br = bufio.NewReader(conn)
	c.bw = bufio.NewWriter(conn)
	return c
}

func (c *Conn) Do(cmd string, argv ...Argv) (interface{}, error) {
	var (
		d                     interface{}
		err                   error
		cmd_buf               bytes.Buffer
		argument              Argv
		cmd_res_protocol_func cmdResProtocolFunc
	)

	// Set argument default values
	if len(argv) > 0 {
		argument = argv[0]
	}

	if argument.Pri == 0 {
		argument.Pri = def_pri
	}
	if len(argument.Tube) == 0 {
		argument.Tube = def_tube
	}

	// Client cmd check
	if row, exists := cmdMaps[cmd]; exists {
		cmd_buf = row.cmd_req_protocol_func(argument)
		cmd_res_protocol_func = row.cmd_res_protocol_func
	} else {
		panic(fmt.Errorf("ErrorCmd :%s does not support", cmd))
	}
	if _, err = cmd_buf.WriteTo(c.bw); err != nil {
		panic(fmt.Errorf("ErrorWrite :%s %s", cmd, err))
	}
	if err = flush(c.bw); err != nil {
		panic(fmt.Errorf("ErrorFlush :%s %s", cmd, err))
	}

	// Process response
	response_argv := cmd_res_protocol_func()
	d, err = response(c.br, response_argv)
	c.br.Reset(c.conn)
	return d, err
}

// Parse response to buffer
func response(r *bufio.Reader, response_argv Response) (interface{}, error) {
	var (
		index    uint8
		data     [][]byte
		body_len int64
	)
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			panic(fmt.Errorf("ErrorResponse :%s", err))
		}
		length := len(line)
		if length == 0 {
			panic(fmt.Errorf("ErrorResponse :%s", "read 0 bytes"))
		}
		eof, err := parse_response(line, index, response_argv, &data, &body_len)
		if err != nil {
			return nil, err
		}
		if eof {
			break
		}
		index++
	}
	return response_argv.parser(data)
}

func parse_response(line []byte, index uint8, response_argv Response, data *[][]byte, body_len *int64) (bool, error) {
	line_length := len(line)

	// Check line suffix
	line_length--
	if line[line_length] != '\n' {
		panic(fmt.Errorf("ErrorResponse :%s", "protocol error"))
	}

	// Rewrite line
	line = line[:line_length]
	if index == 0 {
		// Check first line suffix
		line_length--
		if line[line_length] != '\r' {
			panic(fmt.Errorf("ErrorResponse :%s", "protocol error"))
		}

		// Rwrite line
		line = line[:line_length]

		// Match status
		if ok := bytes.HasPrefix(line, []byte(response_argv.flag)); !ok {
			return true, fmt.Errorf("ErrorResponseParse :%s", line)
		}
		if line_length > len(response_argv.flag)+1 {
			*data = append(*data, line[len(response_argv.flag)+1:])
		}
		if response_argv.multi == false {
			return true, nil
		} else {
			*body_len = parseBodyLen(line[len(response_argv.flag)+1:])
		}
	} else {
		// Write message body to buffer
		*data = append(*data, line)
		*body_len -= int64(line_length) + 1
		if *body_len <= 0 {
			return true, nil
		}
	}
	return false, nil
}

// Parse body length
func parseBodyLen(head []byte) int64 {
	len_buf := getBytes(head, 32)
	length, err := bytesToInt(len_buf)
	if err != nil {
		panic(fmt.Errorf("ErrorParseHead :%s", err))
	}
	return length.(int64)
}

// Bytes to int
func bytesToInt(p []byte) (interface{}, error) {
	if len(p) == 0 {
		return 0, fmt.Errorf("malformed integer")
	}

	var negate bool
	if p[0] == '-' {
		negate = true
		p = p[1:]
		if len(p) == 0 {
			return 0, fmt.Errorf("malformed integer")
		}
	}

	var n int64
	for _, b := range p {
		n *= 10
		if b < '0' || b > '9' {
			return 0, fmt.Errorf("illegal bytes in length")
		}
		n += int64(b - '0')
	}

	if negate {
		n = -n
	}
	return n, nil
}

func getBytes(head []byte, c byte) []byte {
	index := bytes.LastIndexByte(head, c)
	if index == -1 {
		return head
	}
	index++
	if len(head) < index {
		return nil
	}
	return head[index:]
}

// Flush buffer to server
func flush(w *bufio.Writer) error {
	if w.Buffered() > 0 {
		return w.Flush()
	}
	return fmt.Errorf("ErrorFlush %s:", "empty buffer")
}
