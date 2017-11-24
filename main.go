package main

import (
	be "beanstalkd"
	_ "bufio"
	"fmt"
)

func main() {

	//_ = be.Dial
	c := be.Dial("tcp", "127.0.0.1:11300")
	//	id, err := c.Do("put", be.Argv{Message: "this is test"}, be.Argv{})
	//	fmt.Println(id.(int64), err)

	//	c.Do("put")
	//row, _ := c.Do("stats")
	id, err := c.Do("use", be.Argv{Tube: "default"})
	//fmt.Println(id, err)

	//id, err := c.Do("use", be.Argv{Tube: "a_x"})
	//fmt.Println(id, err)

	//id, err = c.Do("delete", be.Argv{Id: 13})
	fmt.Println(id, err)
	/*
		xid := id.(map[string]interface{})["id"].(int64)
		id, err = c.Do("release", be.Argv{Id: xid, Pri: 1024, Delay: 1})
		fmt.Println(id, err)
	*/

}
