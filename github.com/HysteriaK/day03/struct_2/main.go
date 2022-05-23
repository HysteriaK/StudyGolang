package main

import "fmt"

//Address结构体
type Address struct{
	province string
	city string
}

type User struct{
	name string
	gender string
	address Address
}


func main(){
	user := User{
		name : "mike",
		gender: "male",
		address: Address{
			province: "jiangsu",
			city: "nantong",
		},
	}
	fmt.Println(user)
}