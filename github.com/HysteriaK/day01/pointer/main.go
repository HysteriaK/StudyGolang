package main

import "fmt"

func main() {
	//1. &:取地址
	n := 18
	p := &n
	fmt.Println(p) //0xc000014098
	//输出 看看P是什么类型的
	fmt.Printf("%T\n",p) //*int *代表它这是一个指针

	//2. *:根据地址取值
	m := *p
	fmt.Println(m) //18
	fmt.Printf("%T\n",m)

	var a1 *int
	fmt.Println(a1) // nil
	var a2 = new(int)
	fmt.Println(a2) // 0xc0000140e8
	fmt.Println(*a2) // 0
	*a2 = 100
	fmt.Println(*a2) // 100


}