package main

import (
	"fmt"
	"sync"
	"time"
)

//读写互斥锁

var(
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
	rwLock sync.RWMutex //读写锁
)

func read(){
	lock.Lock()
	time.Sleep(time.Millisecond)
	lock.Unlock()
	wg.Done()
}

func write(){
	lock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond*10)
	lock.Unlock()
	wg.Done()
}

func readWithLock(){
	rwLock.RLock()//加读锁
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()//解除读锁
	wg.Done()
}

func writeWithLock(){
	rwLock.Lock()//加写锁
	x = x + 1
	time.Sleep(time.Millisecond*10)
	rwLock.Unlock()//接触读锁
	wg.Done()
}

func main(){
	start := time.Now()

	for i := 0; i < 1000; i++{
		wg.Add(1)
		//go read()
		go readWithLock()
	}

	for i := 0; i < 10; i++{
		wg.Add(1)
		//go write()
		go writeWithLock()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}