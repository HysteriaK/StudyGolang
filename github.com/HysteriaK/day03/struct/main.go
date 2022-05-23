package main

import "fmt"

type person struct{
	name string
	age int
	city string
}

func main(){
	var p1 person
	p1.name = "Mike"
	p1.age = 18
	p1.city = "shanghai"
	fmt.Println(p1)
	fmt.Printf("p1=%v\n",p1)
	var user struct{name string;age int}
	user.name = "Lisa"
	user.age = 20
	fmt.Println(user)
	p5 := person{
		name: "Julia",
		age : 25,
		city : "Beijing",
	}

	fmt.Println(p5)
}