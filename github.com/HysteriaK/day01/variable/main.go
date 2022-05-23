package main

import "fmt"

//声明变量
// var name string
// var age int
// var isOk bool

//批量声明
// var (
// 	name string
// 	age  int
// 	isOk bool
// )

func test() (string, int) {
	return "jrs", 10
}

func main() {

	// name = "leo"
	// age = 16
	// isOk = true
	// fmt.Print(isOk)
	// fmt.Println()
	// fmt.Printf("name:%s", name)
	// fmt.Println(age)

	//声明变量同时赋值
	var s1 string = "julia"
	fmt.Println(s1)
	//类型推导，它可以根据等号右边的内容，来推导变量的类型完成初始化
	var s2 = "julia"
	fmt.Println(s2)
	var s3 = 20
	fmt.Println(s3)

	//短变量声明（只能在函数中使用）
	s4 := "kim"
	fmt.Println(s4)

	//匿名变量
	x, _ := test()
	_, y := test()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
