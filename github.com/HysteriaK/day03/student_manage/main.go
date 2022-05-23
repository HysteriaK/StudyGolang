package main

import (
	"fmt"
	"os"
	"sort"
)

var(
	allStudent map[int]*student
)


type student struct{
	id int
	name string
}

//student的构造函数
func newStudent(id int, name string) *student{
	return &student{
		id: id,
		name: name,
	}
}

func showAllStudent(){
	//如果按照学号顺序来打印所有学生
	//用切片keys存储所有的学生id
	var keys = make([]int,0,200)
	for key := range allStudent{
		keys = append(keys,key)
	}
	//给学生id排序
	sort.Ints(keys)
	//遍历排序过的切片，再根据切片的id顺序打印学生map
	for _,v := range keys{
		fmt.Println(v,allStudent[v].name)
	}
}

func addNewStudent(){
	fmt.Print("请输入学生的id:")
	var id int
	fmt.Scanln(&id)
	fmt.Print("请输入学生的name:")
	var name string
	fmt.Scanln(&name)
	//判断学号是否存在
	_,ok := allStudent[id]
	if !ok{
		stu := newStudent(id,name)
		allStudent[id] = stu
	}else{
		fmt.Println("该学号已经存在，请换个学号")
	}

}

func deleteStudent(){
	fmt.Println("请输入你要删除学生的学号")
	var id int
	fmt.Scan(&id)
	_,ok := allStudent[id]
	if !ok{
		fmt.Println("不存在该学生")
	}else{
		delete(allStudent,id)
	}
}


func main(){
	allStudent = make(map[int]*student,48) //map需要初始化，开辟空间
	for{
			//1.打印菜单
	fmt.Println("学生管理系统")
	fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.退出
	`)
	fmt.Print("请输入你的选择：")
	//2.等待用户选择
	var choice int
	fmt.Scanln(&choice)
	fmt.Printf("你选择了%d这个选项\n",choice)
	//3.执行选择
	switch choice{
	case 1:
		showAllStudent()
	case 2:
		addNewStudent()
	case 3:
		deleteStudent()
	case 4:
		os.Exit(1)
	default:
		fmt.Println("错误输入")
	}

	}
}