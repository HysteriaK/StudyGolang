package main

import (
	"encoding/json"
	"fmt"
)

//结构体和json
//1.序列化：把Go语言中的结构体 ---> json格式的字符串
//2.反序列化：json格式的字符串 ---> Go语言中能够识别的结构体变量

type person struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func main(){
	p := person{
		Name : "Mike",
		Age: 20,
	}
	//序列化，把Go语言中的结构体变成JSON格式的字符串
	//Marshal返回的是切片和一个error
	b, err := json.Marshal(p)
	if err != nil{
		fmt.Printf("marshel failed, err:%v", err)
		return
	}
	fmt.Printf("%#v\n",string(b))

	//反序列化
	str := `{"name":"Mike","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //注意这里是传指针，不然仅仅是拷贝一个副本值，只在函数内部有用
	fmt.Printf("%#v\n",p2)
}