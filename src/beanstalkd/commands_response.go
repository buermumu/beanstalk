package beanstalkd

import (
	"bytes"
	"fmt"
)

type Response struct {
	flag   string
	multi  bool
	parser func([][]byte) (interface{}, error)
}

func cmd_res_put() Response {
	_ = fmt.Println
	return Response{
		flag:   "INSERTED",
		parser: parser_use_first_int,
	}
}

func cmd_res_use() Response {
	return Response{
		flag: "USING",
		parser: func(res [][]byte) (interface{}, error) {
			d := res[0]
			return string(d), nil
		},
	}
}

// Reserve job from current tube
func cmd_res_reserve() Response {
	return Response{
		flag:   "RESERVED",
		multi:  true,
		parser: parser_use_id_body,
	}
}

func cmd_res_delete() Response {
	return Response{
		flag:   "DELETED",
		parser: parser_use_zero_bytes,
	}
}

func cmd_res_release() Response {
	return Response{
		flag:   "RELEASED",
		parser: parser_use_zero_bytes,
	}
}

func cmd_res_bury() Response {
	return Response{
		flag:   "BURIED",
		parser: parser_use_zero_bytes,
	}
}

func cmd_res_touch() Response {
	return Response{
		flag:   "TOUCHED",
		parser: parser_use_zero_bytes,
	}
}

func cmd_res_watch() Response {
	return Response{
		flag:   "WATCHING",
		parser: parser_use_first_int,
	}
}

func cmd_res_ignore() Response {
	return Response{
		flag:   "WATCHING",
		parser: parser_use_first_int,
	}
}

func cmd_res_peek() Response {
	return Response{
		flag:   "FOUND",
		multi:  true,
		parser: parser_use_id_body,
	}
}

func cmd_res_peek_ready() Response {
	return Response{
		flag:   "FOUND",
		multi:  true,
		parser: parser_use_id_body,
	}
}

func cmd_res_peek_delayed() Response {
	return Response{
		flag:   "FOUND",
		multi:  true,
		parser: parser_use_id_body,
	}
}

func cmd_res_peek_buried() Response {
	return Response{
		flag:   "FOUND",
		multi:  true,
		parser: parser_use_id_body,
	}
}

func cmd_res_kick() Response {
	return Response{
		flag:   "KICKED",
		parser: parser_use_first_int,
	}
}

func cmd_res_stats_job() Response {
	return Response{
		flag:   "OK",
		multi:  true,
		parser: parser_use_map,
	}
}

func cmd_res_stats_tube() Response {
	return Response{
		flag:   "OK",
		multi:  true,
		parser: parser_use_map,
	}
}

func cmd_res_stats() Response {
	return Response{
		flag:   "OK",
		multi:  true,
		parser: parser_use_map,
	}
}

// todo
func cmd_res_list_tubes() Response {
	return Response{
		flag:  "OK",
		multi: true,
		parser: func(res [][]byte) (interface{}, error) {
			var tubes []string
			res = res[2:]
			for _, item := range res {
				tubes = append(tubes, string(item[2:]))
			}
			return tubes, nil
		},
	}
}

func cmd_res_list_tube_used() Response {
	return Response{
		flag: "USING",
		parser: func(res [][]byte) (interface{}, error) {
			d := res[0]
			return string(d), nil
		},
	}
}

// todo
func cmd_res_list_tubes_watched() Response {
	return Response{
		flag:  "OK",
		multi: true,
		parser: func(res [][]byte) (interface{}, error) {
			var tubes []string
			res = res[2:]
			for _, item := range res {
				tubes = append(tubes, string(item[2:]))
			}
			return tubes, nil
		},
	}
}

func cmd_res_pause_tube() Response {
	return Response{
		flag:   "PAUSED",
		parser: parser_use_zero_bytes,
	}
}

// Parsed map
func parser_use_map(res [][]byte) (interface{}, error) {
	d := make(map[string]interface{})
	for _, item := range res {
		row := bytes.Split(item, []byte(":"))
		if len(row) < 2 {
			continue
		}
		d[string(row[0])] = row[1]
	}
	return d, nil
}

// Zero is ok
func parser_use_zero_bytes(res [][]byte) (interface{}, error) {
	return len(res) == 0, nil
}

// Parse first element
func parser_use_first_int(res [][]byte) (interface{}, error) {
	d, err := bytesToInt(res[0])
	if err != nil {
		return nil, err
	}
	return d.(int64), nil
}

// Parse map {id:xx, data:xx}
func parser_use_id_body(res [][]byte) (interface{}, error) {
	d := make(map[string]interface{})
	items := bytes.Split(res[0], []byte(" "))
	id, _ := bytesToInt(items[0])
	data := res[1]
	data = data[:len(data)-1]
	d["id"] = id.(int64)
	d["data"] = string(data)
	return d, nil
}
