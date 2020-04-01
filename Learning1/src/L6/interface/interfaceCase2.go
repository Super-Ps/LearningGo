//  接口继承
// 接口是一个引用类型，如果没有对interface 初始化就使用，那么会输出 nil
// 空接口 interfeca{} 没有任何方法， 所以，所有类型都实现了空接口，既我们可以把任何变量赋值给 空接口
//  接口内嵌套别的 接口，有相同的方法时，编译错误，因为一个接口内有重复的方法名

package main

import "fmt"

// 一个接口（比如A 接口）可以继承多个别的接口 （比如B,C），这时如果要实现A接口，也必须将B，C接口的方法也全部实现

//Ainterface  A接口
type Ainterface interface {
	Binterface
	Cinterface
	testA()
}

//Binterface  B接口
type Binterface interface {
	testB()
}

//Cinterface  C接口
type Cinterface interface {
	testC()
}

// 如果要实现 Ainterface 接口，就需要将 Binterface Cinterface 的方法都实现

type stu struct {
}

func (s stu) testA() {
	fmt.Println("this is testA")
}
func (s stu) testB() {
	fmt.Println("this is testB")
}

func (s stu) testC() {
	fmt.Println("this is testC")
}

//T 空接口
type T interface {
}

func main() {

	s := stu{}
	var a Ainterface = s
	a.testA()
	s.testA()

	s2 := stu{}
	var t T = s2
	fmt.Println("空接口的值", t)


	

}
