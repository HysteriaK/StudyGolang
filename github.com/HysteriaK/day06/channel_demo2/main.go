package main

import "fmt"

/*
两个goroutine，两个channel
	1.生成0-100的数字发送到ch1
	2.从ch1中取出数据计算他的平方，把结果发送到ch2中
*/

//生成0-100的数字发送到ch1
func f1(ch chan int){
	for i := 0; i < 100; i++{
		ch <- i
	}
	close(ch)
}

//从ch1中取出数据计算他的平方，把结果发送到ch2中
func f2(ch1,ch2 chan int){
	//如何从通道中循环取值？
	for{
		//如果ok为false，说明ch1通道中的值取完了，可以关闭了
		tmp, ok := <- ch1
		//如果为false说明值取完了,ch1关闭了，退出循环
		if !ok{
			break
		}
		//否则就循环的把tmp值加入通道ch2中
		ch2 <- tmp*tmp
	}
	//当退出循环的时候说明值取完了，也添加完了，关闭通道2
	close(ch2)
}


func main(){
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)
	go f2(ch1,ch2)

	//从ch2中取值 通道中循环取值的方式2
	for ret := range ch2{
		fmt.Println(ret)
	}

}