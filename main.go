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
	//id, err := c.Do("use", be.Argv{Tube: "default"})
	//fmt.Println(id, err)

	id, err := c.Do("watch", be.Argv{Tube: "test"})
	fmt.Println(id, err)

	//id, err = c.Do("delete", be.Argv{Id: 7})
	//fmt.Println(id, err)
	/*
		xid := id.(map[string]interface{})["id"].(int64)
		id, err = c.Do("release", be.Argv{Id: xid, Pri: 1024, Delay: 1})
		fmt.Println(id, err)
	*/

	id, err = c.Do("list-tubes-watched")
	fmt.Println(id, err)

	id, err = c.Do("list-tubes")
	fmt.Println(id, err)
}

func t() interface{} {
	cmd := []interface{}{
		1,
		2,
	}
	return cmd
}

func ternary(expression int, a, b interface{}) interface{} {
	if expression > 0 {
		return a
	}
	return b

}

func x(a *[]interface{}) {
	fmt.Printf("a:%p", *a)
	*a = append(*a, 1)
	*a = append(*a, "xx")
	*a = append(*a, "xyyyx")
}

func parseProtocol(response string) {
}
