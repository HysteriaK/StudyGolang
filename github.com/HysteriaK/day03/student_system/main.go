package main

import (
	"fmt"
	"os"
	"sort"
)

type student struct {
	id   int
	name string
}

func newStudent(id int, name string) *student{
	return &student{
		id:id,
		name:name,
	}
}

type class struct{
	Map map[int]*student
}

func(c class) showAllStudent(){
	var keys []int = make([]int,0,200)
	for key := range c.Map{
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _,v := range keys{
		fmt.Println(v,c.Map[v].name)
	}
	
}

func(c class) addNewStudent(){
	fmt.Print("请输入学生的id:")
	var id int
	fmt.Scanln(&id)
	fmt.Print("请输入学生的name:")
	var name string
	fmt.Scanln(&name)
	//判断输入学生的id是否已经存在
	_,ok := c.Map[id]
	if !ok{
		stu :=newStudent(id,name)
		c.Map[id] = stu
	}else{
		fmt.Println("该学号已经存在，请更改学号后重新输入")
	}

}

func(c class) deleteStudent(){
	fmt.Println("请输入你要删除的学生的学号")
	var id int
	fmt.Scanln(&id)

	//判断该学生id是否存在
	_,ok := c.Map[id]
	if !ok{
		fmt.Println("不存在该学生，请重新输入")
	}else{
		delete(c.Map,id)
	}
}

func(c class) updateStudent(){
	fmt.Println("请输入你要修改的学生id")
	var id int
	fmt.Scanln(&id)

	_,ok := c.Map[id]
	//如果存在这个id
	if ok{
		delete(c.Map,id)
		fmt.Print("请输入更新后的学生id:")
		var newId int
		fmt.Scanln(&newId)
		fmt.Print("请输入更新后的学生name:")
		var newName string
		fmt.Scanln(&newName)
		newStu := newStudent(newId,newName)
		c.Map[newId] = newStu
	}else{
		fmt.Println("你要修改的学生不存在，请重新输入！")
	}
}

func main() {
	c := class{}
	c.Map = make(map[int]*student,200)
	for {
		//打印菜单
		fmt.Println("学生管理系统")
		fmt.Println(`
			1.查看所有学生
			2.添加学生信息
			3.删除学生信息
			4.修改学生信息
			5.退出系统
		`)
		fmt.Println("请输入你的选择：")
		var choice int
		fmt.Scanln(&choice)

		switch choice{
		case 1:
			c.showAllStudent()
		case 2:
			c.addNewStudent()
		case 3:
			c.deleteStudent()
		case 4:
			c.updateStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("操作无效")
		}
	}
}