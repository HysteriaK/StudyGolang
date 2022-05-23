package main

import "fmt"

//常量
//const pi = 3.14159
const (
	STATUS = 200
	NOT_FOUND = 404
)

//批量声明多个常量，如果某一行声明后没有赋值，就默认和上一行相同
// const (
// 	n1 = 100
// 	n2
// 	n3
// )

//iota
// const (
// 	a1 = iota
// 	a2
// 	_
// 	a4
// )

//插队
// const (
// 	c1 = iota //0
// 	c2 = 100  //100
// 	c3 = iota
// 	c4
// )

//iota在同一行
const (
	d1,d2 = iota + 1, iota + 2 //d1:1 d2:2 这一行iota都为0
	d3,d4 = iota + 1, iota + 2 //d3:2 d4:3 这一行iota都为1，因为每增加一行常量声明，iota计数都会加一
)

func main(){
	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(a4)

	// fmt.Println(c1)
	// fmt.Println(c2)
	// fmt.Println(c3)
	// fmt.Println(c4)
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
	
}
