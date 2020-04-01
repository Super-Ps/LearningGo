package main

import "fmt"

//OldMonkeny 被继承
type OldMonkeny struct {
	Name string
}

//NewMonkeny  嵌套继承
type NewMonkeny struct {
	OldMonkeny
}

func (o OldMonkeny) pashu() {
	fmt.Println(o.Name, "天生会爬树")
}

//Fly 让NewMonkeny实现flying
type Fly interface {
	flying()
}

//Swi 让NewMonkeny实现Swiming
type Swi interface {
	Swiming()
}

//flying 实现接口
func (n NewMonkeny) flying() {

	fmt.Println(n.Name, "新增功能会飞啦")
}

//Swiming 实现接口
func (n NewMonkeny) Swiming() {

	fmt.Println(n.Name, "新增功能会游泳了啦")
}

func main() {

	n := NewMonkeny{
		OldMonkeny{
			Name: "原生猴",
		},
	}

	n.pashu()
	n.flying()
	n.Swiming()
}

// 总结：
/*
1.当A 结构体继承了B结构体，那么 A 结构体就自动的集成了 B 结构体的字段和方法，并且可以直接使用
2. 当A 结构体 需要扩展功能，同时不希望去破坏继承关系，则可以去实现某个接口即可
因为我们可以认为：实现接口是对继承机制的补充


接口的继承解决的问题不同：
继承（嵌套）的价值在于： 解决代码的复用性和可维护性
接口的价值在于： 设计，设计好各种规范（方法），让其他自定义类去实现这些方法，接口在一定程度上实现代码解耦
*/
