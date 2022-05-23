package main

import "fmt"

func main(){
	var i1 = 101
	fmt.Printf("%d\n", i1)
	i2 := 077
	fmt.Printf("%o\n", i2)
	fmt.Printf("%d\n", i2)
	i3 := 0xff
	fmt.Printf("%x\n", i3)
	fmt.Printf("%X\n", i3)
	fmt.Printf("%d\n", i3)

	i4 := int16(20) 
	fmt.Printf("%T\n", i4)
}