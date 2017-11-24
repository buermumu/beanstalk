package beanstalkd

import (
	"fmt"
)

func assert_cmd_use(cmd []interface{}) error {
	tube := cmd[1].(string)
	if true {
		// todo assert name char
	}
	if len(tube) > int(max_tube) {
		return fmt.Errorf("assert fail tube:%s more than %d bytes", tube, max_tube)
	}
	return nil
}
