package main

import "fmt"

func main(){
	f1 := 1.23456
	fmt.Printf("%T\n",f1) //默认go语言中的小数都是float64类型

	f2 := float32(1.23456)
	fmt.Printf("%T\n",f2)
	f1 = float64(f2)  //float64和float32的转换
	fmt.Printf("%T\n",f1)

}