package main

import "fmt"

//接口示例2
//不管是什么牌子车都能跑

type car interface{
	run()
}

type Telsa struct{}

func(t Telsa) run(){
	fmt.Println("特斯拉是电车")
}

type Audi struct{}

func(a Audi) run(){
	fmt.Println("奥迪油电混合")
}


func drive(c car){
	c.run()
}

func main(){
	var audi Audi
	var telsa Telsa
	drive(audi)
	drive(telsa)
}