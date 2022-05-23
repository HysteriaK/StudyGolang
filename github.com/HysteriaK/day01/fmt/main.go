package main

import "fmt"

//fmt占位符
func main(){
	var n = 100
	//查看类型 %T
	fmt.Printf("%T\n", n)
	//查看变量的value
	fmt.Printf("%v\n", n)

	//查看十进制的值
	fmt.Printf("%d\n", n)
	//查看二进制的值
	fmt.Printf("%b\n", n)
	//查看八进制的值
	fmt.Printf("%o\n", n)
	//查看十六进制的值
	fmt.Printf("%x\n", n)

	//查看字符串
	var s = "hello"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
}