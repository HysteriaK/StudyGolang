package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string)string{
	return func(name string)string{
		if !strings.HasSuffix(name,suffix){
			return name + suffix
		}
		return name
	}
}

func main(){
	ret := makeSuffixFunc(".jpg")
	fmt.Println(ret("test"))
}