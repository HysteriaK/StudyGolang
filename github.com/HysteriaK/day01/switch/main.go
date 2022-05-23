package main

import "fmt"

//switch
//简化大量的判断
func main(){
	// var n = 5
	// if n == 1{
	// 	fmt.Println("大拇指")
	// }else if n==2{
	// 	fmt.Println("食指")
	// }else if n==3{
	// 	fmt.Println("中指")
	// }else if n==4{
	// 	fmt.Println("无名指")
	// }else if n==5{
	// 	fmt.Println("小拇指")
	// }else{
	// 	fmt.Println("无效数字")
	// }

	//switch简化
	//要么在外面申明n的值
	//例如 var n = 3
	//要么在switch里申明
	switch n:=3;n{
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效数值")
	}

	age := 20
	switch{
	case age < 18:
		fmt.Println("未成年")
	case age >= 18 && age <=30:
		fmt.Println("青年")
	case age >30 && age < 60:
		fmt.Println("壮年")
	case age >= 70:
		fmt.Println("老年")
	default:
		fmt.Println("不是人")
	}
}