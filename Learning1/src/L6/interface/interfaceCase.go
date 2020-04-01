package main

import "fmt"

//Ainterface test
type  Ainterface interface{
	Say()
}


type stu struct{
	Name string 
}

func (s stu) Say(){
	fmt.Println("stu say")
}



//只要是自定义数据类型，就可以实现接口 不仅仅是结构体类型。
type integer  int


func (i integer) Say(){
	fmt.Println("stu say",i)
}


func main(){

	var a  Ainterface  // 1》接口本身不能创建实例 是调不起来里面的方法，因为该方法还没实现
	// a.Say()   // 错误 
	var s stu   // 结构体变量，实现了say，实现了  Ainterface，所以 a.Say()能运行成功
	a = s  // 2》但是可以指向一个实现了该接口的自定义类型的变量(实例)
	a.Say()
	
	var i integer =10000
	var b  Ainterface = i
	b.Say()
}