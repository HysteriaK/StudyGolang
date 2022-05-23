package main

import "fmt"

func f1(){
	fmt.Println("hello")
}

func f2() int{
	return 10
}


//函数也可以作为参数的类型
func f3(x func() int){
	ret := x()
	fmt.Println(ret)
}

func add(x,y int)int{
	return x+y
}

func calc(x,y int, op func(int,int)int)int{
	return op(x,y)
}


func main(){

	a := f1
	fmt.Printf("%T\n",a)  //func()
	b := f2
	fmt.Printf("%T\n",b) //func() int

	f3(f2)

	ret := calc(2,3,add)
	fmt.Println(ret)


}