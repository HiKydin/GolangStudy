package main

import "fmt"

func df1() {
	fmt.Println("A")
}
func df2() {
	fmt.Println("B")
}
func df3() {
	fmt.Println("C")
}
func main() {
	defer df1()
	defer df2()
	defer df3()
}

//defer会在当前函数结束前运行，defer执行顺序为压栈
//defer比return后执行
