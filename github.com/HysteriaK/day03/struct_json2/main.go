package main

import (
	"encoding/json"
	"fmt"
)

type Student struct{
	Id int
	Gender string
	Name string
}

type Class struct{
	Title string
	Students []*Student
}


func main(){
	c := &Class{
		Title: "101",
		Students: make([]*Student, 0, 100),
	}
	for i := 0; i < 10; i++{
		stu := &Student{
			Name: fmt.Sprintf("stu%02d\n",i),
			Gender: "男",
			Id: i,
		}
		c.Students = append(c.Students, stu)
	}

	//json序列化
	data, err := json.Marshal(c)
	if err != nil{
		fmt.Println("marshal failed")
		return
	}
	fmt.Printf("json:%s\n",data)


}