package main

import "fmt"

//panic和recover

func funcA(){
	fmt.Println("a")
}

func funcB(){

	defer func(){
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接...")
	}()
	panic("严重出错") //程序崩溃退出，打印b，以及函数funcC都不会再执行

	fmt.Println("b")
}

func funcC(){
	fmt.Println("c")
}

func main(){
	funcA()
	funcB()
	funcC()
}