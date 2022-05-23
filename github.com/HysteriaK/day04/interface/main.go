package main

import "fmt"

//引出接口的实例

type cat struct{}

type dog struct{}

func (c cat) Speak() {
	fmt.Println("miao")
}

func (d dog) Speak(){
	fmt.Println("wang")
}

//想要这个beat方法能传任何参数，可以是猫可以是狗也可以是别的动物
//我们就可以用接口来实现，我们可以实现一个Speaker类型，它必须实现Speak方法，只要被打了我们就调用speak方法
type Speaker interface{
	Speak()
}

func beat(s Speaker){
	s.Speak()
}

func main(){
	var c1 cat
	beat(c1)
	var d1 dog
	beat(d1)
}