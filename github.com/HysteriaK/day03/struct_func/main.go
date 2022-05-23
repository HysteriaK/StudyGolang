package main

import "fmt"

//结构体的构造函数

type person struct{
	name string
	age int
}

func newPerson(name string, age int) person{
	return person{
		name: name,
		age : age,
	}
}

//指针类型接收者
func(p *person) setAge(newAge int){
	p.age = newAge
}

//值类型接收者
func(p person) setAge2(newAge int){
	p.age = newAge
}

func main(){
	p := newPerson("Echo",20)
	fmt.Println(p.age)
	p.setAge(25)
	fmt.Println(p.age)
	p.setAge2(30)
	fmt.Println(p.age)
}