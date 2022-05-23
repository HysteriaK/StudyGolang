package main

import (
	"fmt"
	"sync"
)

//多个goroutine并发操作全局变量x

var(
	x int64
	wg sync.WaitGroup
	lock sync.Mutex //互斥锁，只有一个人能够获得这个锁
)

func add(){
	for i := 0; i < 5000; i++{
		lock.Lock()//加锁，别的goroutine就拿不到了
		x = x + 1
		lock.Unlock()//做完操作后就释放锁
	}
	wg.Done()
}

func main(){
	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}