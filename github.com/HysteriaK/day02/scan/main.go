package main

import "fmt"

func main(){
	var s string
	fmt.Scan(&s)
	fmt.Println(s)

	var(
		name string
		age int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	fmt.Printf("输入的结果是 name:%s age:%d married:%t \n",name,age,married)
}