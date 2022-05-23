package main

import "fmt"

type dog struct {
	name string
	kind string
	age  int
}

//构造方法
func newDog(name, kind string, age int) *dog {
	return &dog{
		name: name,
		kind: kind,
		age:  age,
	}
}

//方法是作用于特别类型的函数
//特定类型变量狗的方法
func (d dog) bark() {
	fmt.Println("狗叫")
}

func main() {
	p := newDog("java","samo",1)
	p.bark()
}