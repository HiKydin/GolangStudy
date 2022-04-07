package main

import (
	"fmt"
	_ "strconv"
)

var (
	d bool = true
	e int  = 666
)

func main() {
	var a int = 10
	var b = 20
	var c = "hello world"
	fmt.Println(a, "\n", b, "\r", c, "\n", d, "\n", e)
	var f = 10
	fmt.Printf("f的类型是 %T ", f)

	// var str string = "true"
	// var b bool
	// b, _ = strconv.ParseBool(str)
	fmt.Println(b)
}
