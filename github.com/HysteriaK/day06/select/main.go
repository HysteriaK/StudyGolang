package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
			//分支一：如果能从ch中取到值
			case x := <-ch:
				fmt.Println(x)
			//分支二：如果能在ch中添加值
			case ch <-i:
			default:
				fmt.Println("不做任何事")
		}
	}
}