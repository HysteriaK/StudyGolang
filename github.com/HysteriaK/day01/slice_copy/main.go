package main

import (
	"fmt"
	"sort"
)

func main(){
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)

	var b = [...]int{3,7,8,9,1}
	sort.Ints(b[:])
	fmt.Println(b)
	
	a1 := [...]int{1,3,5,7,9,11,13,15,17}
	s1 := a1[:]
	//删掉索引为1的那个3
	s1 = append(s1[:1],s1[2:]...)
	fmt.Println(s1)
	fmt.Println(a1)
}