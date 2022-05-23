package main

import (
	"fmt"
	"sync"
)

//sync.map 并发安全的map
var(
	wg sync.WaitGroup
)

var m = make(map[int]int)

//并发安全的map
var m2 = sync.Map{}

func get(key int)int{
	return m[key]
}

func set(key,value int){
	m[key] = value
}

func main(){
	for i := 0; i < 20; i++{
		wg.Add(1)
		go func(i int) {
			m2.Store(i,i+100) //设置键值对
			value, _ := m2.Load(i)
			fmt.Printf("key:%v value:%v", i, value) //打印键值对
			wg.Done()
		}(i)
	}
	wg.Wait()
}