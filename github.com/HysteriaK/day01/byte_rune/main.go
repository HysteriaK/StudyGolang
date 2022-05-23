package main

import "fmt"

func main(){
	s := "Hello嘉兴hotel"
	for i := 0; i < len(s); i++{ //byte
		fmt.Printf("%v(%c)",s[i],s[i])
	}
	fmt.Println()

	for _, r := range s{ //rune
		fmt.Printf("%v(%c)",r,r)
	}
	fmt.Println()

	//可以发现，当字符串里有中文或者其他字符时，我们用rune能表示出来

	//字符串修改
	//byte
	s1 := "big"
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	//将字符数组输出，注意类型转换
	fmt.Println(string(byteS1))

	s2 := "蓝色战衣"
	runes2 := []rune(s2)
	runes2[0] = '红'
	fmt.Println(string(runes2))

	s3 := "hair"
	byteS3 := []byte(s3)
	byteS3[0] = 'f'
	fmt.Println(string(byteS3))
	
	fmt.Println("-----------")

	word := "hello钢铁侠"
	count := 0
	for _, x := range word{
		if(len(string(x)) >= 3){
			count++
		}
	}
	fmt.Println(count)

	s4 := "神奇女侠"
	runeS4 := []rune(s4)
	runeS4[0] = '惊'
	fmt.Println(string(runeS4))









}