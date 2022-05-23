package main

import (
	"fmt"
	"strings"
)

func main(){

	//元素类型为map的切片
	var mapSlice = make([]map[string]string,3)

	for index,value := range mapSlice{
		fmt.Printf("index:%d value:%v\n",index,value)
	}

	//对切片中的元素进行初始化
	mapSlice[0] = make(map[string]string,10)
	mapSlice[0]["name"] = "张三"
	mapSlice[0]["age"] = "18"
	mapSlice[0]["password"] = "123456"

	for index,value := range mapSlice{
		fmt.Printf("index:%d value:%v\n",index,value)
	}

	//值为切片类型的map
	var sliceMap = make(map[string][]string,3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok{
		value = make([]string,0,2)
	}
	value = append(value, "上海","北京")
	sliceMap[key] = value
	fmt.Println(sliceMap)

	str := "how do you do"
	strSlice := strings.Split(str," ")
	strMap := make(map[string]int)
	for _,v := range strSlice{
		_, ok := strMap[v]
		if !ok{
			strMap[v] = 1 
		}else{
			strMap[v] = strMap[v] + 1
		}
	}
	fmt.Println(strMap)




}