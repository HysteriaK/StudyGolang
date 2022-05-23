package main

import "fmt"

func intSum(x int, y int)int{
	return x+y
}

func sayHello(){
	fmt.Println("sayHello")
}

func intSum2(x...int)int{
	fmt.Println(x)
	sum := 0
	for _,v := range x{
		sum = sum + v
	}
	return sum 
}

//多返回值
func calc(x, y int)(int, int){
	sum := x + y
	sub := x - y
	return sum, sub
}

func calc2(x, y int)(sum, sub int){
	sum = x + y
	sub = x - y
	return
}

func main(){
	sayHello()
	res := intSum(2,3)
	fmt.Println(res)
	ret1 := intSum2()
	ret2 := intSum2(10)
	ret3 := intSum2(10,20)
	ret4 := intSum2(10,20,30)
	fmt.Println(ret1,ret2,ret3,ret4)
	fmt.Println(calc(5,2))
	fmt.Println(calc2(6,2))
}