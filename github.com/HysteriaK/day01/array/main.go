package main

import "fmt"

func main(){
	//var testArray [3]int  //数组会初始化为int类型的零值
	//var numArray = [3]int{1,2,3}  //数组会使用指定的值完成初始化
	//var cityArray = [3]string{"北京","上海","深圳"}  //使用指定的值完成初始化

	//var testArray [3]int
	//var numArray = [...]int{1,2}
	var cityArray = [...]string{"深圳","上海","北京","广州"}

	// var a = [7]int{1:2,5:3}
	// fmt.Println(a)

	// fmt.Println(testArray)
	// fmt.Println(numArray)
	// fmt.Printf("Type of numArray:%T\n",numArray)
	// fmt.Println(cityArray)
	// fmt.Printf("Type of cityArray:%T\n",cityArray)

	for i:=0;i<len(cityArray);i++{
		fmt.Println(cityArray[i])
	}

	//for range遍历
	for index,value := range cityArray{
		fmt.Println(index,value)
	}

	// //多维数组
	// var a [3][2]int
	// a = [3][2]int{
	// 	[2]int{1,2},
	// 	[2]int{3,4},
	// 	[2]int{5,6},
	// }
	// fmt.Println(a)

	// //多维数组的遍历
	// for _,v1 := range a{
	// 	for _,v2 := range v1{
	// 		fmt.Println(v2)
	// 	}
	// }

	/*
	1.求数组[1, 3, 5, 7, 8]所有元素的和
	2.找出数组中和为指定值的两个元素的下标
	比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	*/
	b := [...]int{1,3,5,7,8}
	var sum = 0
	var cur = 0
	for i:=0; i<len(b);i++{
		cur = b[i]
		sum = sum + cur
	}
	fmt.Println(sum)
	

	for i,v := range b{
		for j:=i+1;j<len(b);j++{
			if v + b[j] == 8{
				fmt.Println(i,j)
			}
		}
	}

	
}