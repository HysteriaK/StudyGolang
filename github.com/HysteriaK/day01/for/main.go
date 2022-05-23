package main

import "fmt"

func main(){
	//基本的for循环格式
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}

	// 变式一(初始语句可以省略)
	var i = 5
	for ;i<10;i++{
		fmt.Println(i)
	}

	//变种三
	// var i = 5
	// for i < 10{
	// 	fmt.Println(i)
	// 	i++
	// }

	//for range循环
	s := "复仇者联盟avengers"
	for i,v := range s{
		fmt.Printf("%d %c\n", i, v)
	}

	for i:=1;i<=9;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%d * %d = %d\t",i,j,i*j)
		}
		fmt.Println()
	}

	
	//流程控制之跳出for循环
	//当i = 5的时候跳出for循环
	for i:=0;i<10;i++{
		if i == 5{
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")


	//当i = 5的时候，跳过此次for循环，不执行for循环内部的打印语句
	for i:=0;i<10;i++{
		if i==5{
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}