package main

import "fmt"

type mover interface{
	move()
}

type eater interface{
	eat()
}

type cat struct{}

func(c cat) move(){
	fmt.Println("cat can move")
}

func(c cat) eat(){
	fmt.Println("cat can eat")
}

func main(){
	var java cat
	var m mover = java
	var e eater = java
	m.move()
	e.eat()

}