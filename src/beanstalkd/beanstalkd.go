package beanstalkd

import (
	_ "bufio"
	"bytes"
	"fmt"
	"strconv"
)

var (
	// Max priority
	max_pri uint32 = 4294967295

	// Min priority
	min_pri uint32 = 0

	// Default priority
	def_pri uint32 = 1024

	// Default delay second
	def_delay uint64 = 0

	// Default ttr second
	def_ttr int64 = 1

	// Max tube bytes
	max_tube int32 = 200
)

var (
	cmdMaps map[string]cmdExtra
)

type cmdFunc func(Argv) bytes.Buffer
type cmdResFunc func() Response

type cmdExtra struct {
	cmd_func cmdFunc
	cmd_res  cmdResFunc
}

// cmd argvs
type Argv struct {
	Id      int64
	Pri     uint32
	Delay   uint64
	Ttr     uint64
	Timeout int64
	Bound   uint32
	Tube    string
	Message string
}

// cmd untis convert beanstalkd protocol
func convertProtocol(w *bytes.Buffer, cmdUnits []interface{}) (int, error) {
	var buf_space byte = 32
	var buf_r byte = 13
	var buf_n byte = 10
	var size, n int
	var err error
	var dst []byte
	cmd_len := len(cmdUnits)
	for _, item := range cmdUnits {
		switch v := item.(type) {
		case string:
			n, err = w.Write([]byte(string(v)))
		case int:
			n, err = w.Write(strconv.AppendInt(dst, int64(v), 10))
		case int8:
			n, err = w.Write(strconv.AppendInt(dst, int64(v), 10))
		case int16:
			n, err = w.Write(strconv.AppendInt(dst, int64(v), 10))
		case int32:
			n, err = w.Write(strconv.AppendInt(dst, int64(v), 10))
		case int64:
			n, err = w.Write(strconv.AppendInt(dst, int64(v), 10))
		case uint32:
			n, err = w.Write(strconv.AppendInt(dst, int64(v), 10))
		case uint64:
			n, err = w.Write(strconv.AppendInt(dst, int64(v), 10))
		default:
			return size, fmt.Errorf("ErrorProtocol :%s", "Unknow cmd or arg data type")
		}
		if cmd_len > 1 {
			err = w.WriteByte(buf_space)
			n += 1
		}
		cmd_len--
		if err != nil {
			return size, err
		}
		size += n
	}
	err = w.WriteByte(buf_r)
	if err != nil {
		return size, err
	}
	size += 1
	err = w.WriteByte(buf_n)
	if err != nil {
		return size, err
	}
	return size, err
}

func _protocol(cmd []interface{}) bytes.Buffer {
	var buffer bytes.Buffer
	if _, err := convertProtocol(&buffer, cmd); err != nil {
		panic(err)
	}
	return buffer
}

func init() {
	cmdMaps = make(map[string]cmdExtra)
	cmdMaps["put"] = cmdExtra{
		cmd_put,
		cmd_res_put,
	}
	cmdMaps["use"] = cmdExtra{
		cmd_use,
		cmd_res_use,
	}
	cmdMaps["stats"] = cmdExtra{
		cmd_stats,
		cmd_res_stats,
	}
	cmdMaps["reserve"] = cmdExtra{
		cmd_reserve,
		cmd_res_reserve,
	}
	cmdMaps["delete"] = cmdExtra{cmd_delete, cmd_res_delete}
	cmdMaps["release"] = cmdExtra{cmd_release, cmd_res_release}
	cmdMaps["bury"] = cmdExtra{cmd_bury, cmd_res_bury}
	cmdMaps["touch"] = cmdExtra{cmd_touch, cmd_res_touch}
	cmdMaps["watch"] = cmdExtra{cmd_watch, cmd_res_watch}
	cmdMaps["ignore"] = cmdExtra{cmd_ignore, cmd_res_ignore}
	cmdMaps["peek"] = cmdExtra{cmd_peek, cmd_res_peek}
	cmdMaps["peek-ready"] = cmdExtra{cmd_peek_ready, cmd_res_peek_ready}
	cmdMaps["peek-delayed"] = cmdExtra{cmd_peek_delayed, cmd_res_peek_delayed}
	cmdMaps["peek-buried"] = cmdExtra{cmd_peek_buried, cmd_res_peek_buried}
	cmdMaps["kick"] = cmdExtra{cmd_kick, cmd_res_kick}
	cmdMaps["stats-job"] = cmdExtra{cmd_stats_job, cmd_res_stats_job}
	cmdMaps["stats-tube"] = cmdExtra{cmd_stats_tube, cmd_res_stats_tube}
	cmdMaps["list-tubes"] = cmdExtra{cmd_list_tubes, cmd_res_list_tubes}
	cmdMaps["list-tube-used"] = cmdExtra{cmd_list_tube_used, cmd_res_list_tube_used}
	cmdMaps["list-tubes-watched"] = cmdExtra{cmd_list_tubes_watched, cmd_res_list_tubes_watched}
	cmdMaps["pause_tube"] = cmdExtra{cmd_pause_tube, cmd_res_pause_tube}
}
