package main

import "fmt"

//一个接口的所有方法不一定由一个类型全部实现
//接口的方法可以在一个类型中嵌套一个类型或者结构体实现


type WashingMachiner interface{
	wash()
	dry()
}

//甩干器
type dryer struct{
	brand string
}

func(d dryer) dry(){
	fmt.Println("甩干")
}

type Hair struct{
	dryer
}

func(h Hair) wash(){
	fmt.Println("洗衣功能")
}

func main(){
	var w WashingMachiner
	w = Hair{
		dryer{brand: "Hair"},
	}
	w.wash()
	w.dry()


}