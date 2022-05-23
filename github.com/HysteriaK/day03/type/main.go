package main

import "fmt"

//自定义类型和类型别名

//type后面跟的是类型 ----> 自定义类型
type myInt int

//类型别名
type yourInt = int

func main(){
	var n myInt = 100
	var m yourInt = 100
	fmt.Println(n)
	fmt.Printf("%T\n",n)
	fmt.Println(m)
	fmt.Printf("%T\n",m)
}