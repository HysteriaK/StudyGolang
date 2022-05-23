package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main(){
	// scoreMap := make(map[string]int)
	// scoreMap["张三"] = 90
	// scoreMap["李四"] = 80
	// scoreMap["王五"] = 70

	// //遍历map的键值对
	// for k,v := range scoreMap{
	// 	fmt.Println(k,v)
	// }

	// //遍历map的键
	// for k := range scoreMap{
	// 	fmt.Println(k)
	// }

	// delete(scoreMap,"张三")
	// for k,v := range scoreMap{
	// 	fmt.Println(k,v)
	// }

	//按照指定顺序去遍历map


	var gradeMap = make(map[string]int, 200)
	for i := 0; i < 100; i++{
		key := fmt.Sprintf("stu%02d",i) //生成stu开头的字符串
		value := rand.Intn(100) //生成100以内的非负随机数
		gradeMap[key] = value // 把这100个键值对放入map中
	}

	//取出map里的key，放入切片中
	var keys = make([]string,0,200)
	for key := range gradeMap{
		keys = append(keys, key)
	}
	//对切片keys进行排序，相当于就是对map的键进行排序
	sort.Strings(keys)
	//keys切片已经排序过了，那么我们遍历切片，在用遍历出来的一个一个key，去map里找相应的value
	//我们不需要它的索引，所以用_
	for _, key := range keys{
		fmt.Println(key,gradeMap[key])
	}



}