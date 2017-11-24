package beanstalkd

import (
	"bytes"
	_ "fmt"
	_ "time"
)

// Use tube
func cmd_use(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"use",
		argv.Tube,
	}
	if err := assert_tube_name(argv.Tube); err != nil {
		panic(err)
	}
	return _protocol(cmd)
}

// Write job to current tube
func cmd_put(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"put",
		argv.Pri,
		argv.Delay,
		argv.Ttr,
		len(argv.Message),
	}

	var buffer bytes.Buffer
	convertProtocol(&buffer, cmd)
	convertProtocol(&buffer, []interface{}{argv.Message})
	return buffer
}

// Reserve job from current tube
func cmd_reserve(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"reserve",
	}
	// assert_cmd_reserve
	return _protocol(cmd)
}

func cmd_delete(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"delete",
		argv.Id,
	}
	// assert_cmd_delete
	return _protocol(cmd)
}

func cmd_release(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"release",
		argv.Id,
		argv.Pri,
		argv.Delay,
	}
	// assert_cmd_release
	return _protocol(cmd)
}

func cmd_bury(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"bury",
		argv.Id,
		argv.Pri,
	}
	// assert_cmd_bury
	return _protocol(cmd)
}

func cmd_touch(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"touch",
		argv.Id,
	}
	// assert_cmd_touch
	return _protocol(cmd)
}

func cmd_watch(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"watch",
		argv.Tube,
	}
	// assert_cmd_watch
	return _protocol(cmd)
}

func cmd_ignore(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"ignore",
		argv.Tube,
	}
	// assert_cmd_ignore
	return _protocol(cmd)
}

func cmd_peek(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"peek",
		argv.Id,
	}
	// assert_cmd_peek
	return _protocol(cmd)
}

func cmd_peek_ready(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"peek-ready",
	}
	return _protocol(cmd)
}

func cmd_peek_delayed(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"peek-delayed",
	}
	return _protocol(cmd)
}

func cmd_peek_buried(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"peek-buried",
	}
	return _protocol(cmd)
}

func cmd_kick(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"kick",
		argv.Bound,
	}
	return _protocol(cmd)
}

func cmd_stats_job(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"stats-job",
		argv.Id,
	}
	return _protocol(cmd)
}

func cmd_stats_tube(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"stats-tube",
		argv.Tube,
	}
	return _protocol(cmd)
}

func cmd_stats(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"stats",
	}
	return _protocol(cmd)
}

func cmd_list_tubes(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"list-tubes",
	}
	return _protocol(cmd)
}

func cmd_list_tube_used(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"list-tube-used",
	}
	return _protocol(cmd)
}

func cmd_list_tubes_watched(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"list-tubes-watched",
	}
	return _protocol(cmd)
}

func cmd_quit(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"quit",
	}
	return _protocol(cmd)
}

func cmd_pause_tube(argv Argv) bytes.Buffer {
	cmd := []interface{}{
		"pause-tube",
		argv.Tube,
		argv.Delay,
	}
	return _protocol(cmd)
}
