package main

import "fmt"

type Animal struct{
	name string
}

func(a *Animal) move(){
	fmt.Printf("%s会动! \n",a.name)
}

type Dog struct{
	feet int
	*Animal //通过嵌套匿名结构体来实现继承
}

func(d *Dog) wang(){
	fmt.Printf("%s会汪汪汪\n",d.name)
}

func main(){
	d := Dog{
		feet: 14,
		Animal: &Animal{
			name: "Bob",
		},
	}

	d.wang()
	d.move()
}