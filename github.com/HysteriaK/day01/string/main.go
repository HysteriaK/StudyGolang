package main

import (
	"fmt"
	"strings"
)

func main(){

	//字符串的长度
	s2 := "千金散尽还复来"
	fmt.Println(len(s2))
	n := len(s2)
	fmt.Println(n)

	//字符串的拼接
	s1 := "leo"
	s3 := "smart"

	fmt.Println(s1 + s3)
	ss1 := fmt.Sprintf("%s%s",s1,s3) //用Sprintf来拼接，但是Sprintf返回的类型是字符串，所以要去用变量接收
	fmt.Println(ss1)

	//字符串的分割
	address := "7JurongEast,32Street,Mayfair,#1405"
	res := strings.Split(address, ",")
	fmt.Println(res)

	//字符串的包含
	fmt.Println(strings.Contains(ss1,"leo"))

	//字符串的前缀
	fmt.Println(strings.HasPrefix(ss1,"leo"))
	//后缀
	fmt.Println(strings.HasSuffix(ss1,"leo"))


	//子串出现的位置
	ss4 := "abcdefdb"
	fmt.Println(strings.Index(ss4,"c"))
	fmt.Println(strings.LastIndex(ss4,"b"))

	//拼接(把字符数组用加号拼接起来)
	fmt.Println(strings.Join(res,"+"))


}