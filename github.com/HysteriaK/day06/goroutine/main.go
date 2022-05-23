package main

import (
	"fmt"
	"time"
)

func show(msg string){
	for i := 1; i < 5; i++{
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main(){
	go show("java")
	go show("golang")
	fmt.Println("end...")
}