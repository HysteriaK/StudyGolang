package main

import "fmt"

func adder() func(int)int{
	var x = 100
	return func(y int)int{
		x += y
		return x
	}
}

func adder2(x int) func(int)int{
	return func(y int)int{
		x += y
		return x
	}
}


func main(){
	ret := adder()
	fmt.Println(ret(20))
	fmt.Println(ret(40))
	res := adder2(100)
	fmt.Println(res(20))
	fmt.Println(res(40))

}