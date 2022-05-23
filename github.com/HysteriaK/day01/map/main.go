package main

import "fmt"

func main(){
	//var m1 map[string]int //光用var声明时，map还没有被初始化，没有在内存中被开辟空间
	var m1 map[string]int = make(map[string]int,10) //预估好map的容量，避免在程序运行期间在动态扩容
	m1["height"] = 180
	m1["weight"] = 150
	m1["age"] = 25

	fmt.Println(m1)

	scoreMap := make(map[string]int,5)
	scoreMap["张三"] = 88
	scoreMap["李四"] = 98
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	fmt.Printf("type of scoreMap: %T\n",scoreMap)
	value,ok := scoreMap["张三"]
	if ok {
		fmt.Println(value)
	}else{
		fmt.Println("查无此人")
	}



}