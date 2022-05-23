package main

import "fmt"

func main(){
	var s1 []int //定义一个int类型的切片
	var s2 []string //定义一个string类型的切片
	fmt.Println(s1,s2)
	fmt.Println(s1 == nil) //true
	fmt.Println(s2 == nil) //true
	//初始化
	s1 = []int{1,2,3}
	s2 = []string{"深圳","北京","上海"}
	fmt.Println(s1 == nil) //false
	fmt.Println(s2 == nil) //false

	//长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d",len(s1),cap(s1))
	fmt.Println()
	fmt.Printf("len(s2):%d cap(s2):%d",len(s2),cap(s2))
	fmt.Println()

	//2.由数组得到的切片
	a1 := [...]int{1,3,5,7,9,11,13}
	a2 := [...]int{55,56,57,58,59}
	c2 := a2[1:]
	fmt.Println(c2)
	d2 := a2[:4]
	fmt.Println(d2)
	fmt.Println(a2[:])


	//左闭右开区间，从第0个元素取，但是不包括第4个元素
	s3 := a1[0:4] // [1 3 5 7]
	fmt.Printf("len(s3):%d cap(s3):%d",len(s3),cap(s3)) //len是4，cap是7
	fmt.Println()
	fmt.Println(s3)
	c := a1[1:]
	fmt.Println(c)
	d := a1[:4]
	fmt.Println(d)

	s3[0] = 10
	fmt.Println(s3)
	fmt.Println(a1)






	


}