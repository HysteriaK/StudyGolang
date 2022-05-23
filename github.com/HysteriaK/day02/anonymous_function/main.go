package main

import "fmt"

//匿名函数

func main(){
	//函数内部无法声明带名字的函数
	// f1 := func(x,y int){
	// 	fmt.Println(x+y)
	// }
	// f1(2,3)

	//只需调用一次的函数，还可以简写成立即执行函数


	func(x,y int){
		fmt.Println(x+y)
		fmt.Println("hello world")
	}(10,20)

}