package beanstalkd

import (
	"bytes"
	"fmt"
)

var tube_name_chars = "\\-+/;.$_()0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func assert_tube_name(tube string) error {
	tube_len := len(tube)
	if tube_len == 0 {
		return fmt.Errorf("assert fail tube name is empty")
	}
	if tube_len > int(max_tube) {
		return fmt.Errorf("assert fail tube:%s more than %d bytes", tube, max_tube)
	}
	for _, c := range tube {
		if index := bytes.IndexByte([]byte(tube_name_chars), byte(c)); index == -1 {
			return fmt.Errorf("assert fail tube has invalid char:%c", c)
		}
	}
	return nil
}
