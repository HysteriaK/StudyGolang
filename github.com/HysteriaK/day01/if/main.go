package main


import "fmt"

// if条件判断
func main(){
	age := 19
	if age > 18 {
		fmt.Println("澳门首家线上赌场开业！")
	}else{
		fmt.Println("完成假期作业！")
	}

	//多个判断条件
	if age > 35 {
		fmt.Println("中年")
	}else if age > 18{
		fmt.Println("青年")
	}else{
		fmt.Println("读书的年纪")
	}

	//作用域
	//此时的age变量只在if条件判断语句中生效
	if age := 19; age > 18{
		fmt.Println("成人")
	}else{
		fmt.Println("未成人")
	}


}