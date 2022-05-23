package main

import "fmt"

//make()函数创造切片
func main(){
	s1 := make([]int,5,10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n",s1,len(s1),cap(s1))
	s2 := make([]int,0,10)
	fmt.Println(len(s2)==0)
	fmt.Println(s2 == nil)

	s3 := make([]int,3)
	s4 := s3
	fmt.Println(s3)
	s4[0] = 100
	fmt.Println(s3)

	//切片的遍历
	//1.索引遍历
	 s := []int{1,2,3}
	 for i:=0;i<len(s);i++{
		 fmt.Println(i,s[i])
	 }
	 //2.for range遍历
	 for i,v := range s{
		 fmt.Println(i,v)
	 }
}