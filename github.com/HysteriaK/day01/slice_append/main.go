package main

import "fmt"

func main(){
	var s []int
	fmt.Println(s) // []
	s = append(s, 1)
	fmt.Println(s) // [1]
	s = append(s, 2,3,4)
	fmt.Println(s) // [1 2 3 4]
	s2 := []int{5,6,7}
	s = append(s, s2...) // [1 2 3 4 5 6 7]
	fmt.Println(s)
}