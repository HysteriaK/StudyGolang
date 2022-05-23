package main

import "fmt"

func main(){

	//普通的方式退出多层循环
	// var flag = false
	// for i := 0; i < 10; i++{
	// 	for j:='A'; j < 'Z'; j++{
	// 		if j == 'C'{
	// 			flag = true
	// 			break
	// 		}
	// 		fmt.Printf("%v-%c\n",i,j)
	// 	}
	// 	if flag{
	// 		break
	// 	}
	// }
	//使用goto语句退出循环
	for i := 0; i < 10; i++{
		for j := 'A'; j < 'Z'; j++{
			if j == 'C'{
				goto XX
			}
			fmt.Printf("%v-%c\n",i,j)
		}
	}

	XX: //相当于一个标签，然后goto这个标签，就无条件跳到这个标签的位置
	fmt.Println("over")


}