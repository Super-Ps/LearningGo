// 继承：  GO 没有继承，在一个struct中嵌套了另一个匿名的结构体，是通过嵌套，组合的方式
//那么这个结构体可以直接访问匿名结构体的字段和方法，
//从而实现了继承特性
// 没有父类子类的概念， 将共有的存在在一一个结构体内，特有的去集成共有的结构体

package main

import (
	"fmt"
	"log"
	"os"
)

//Acitons 外包可以访问
type Acitons struct {
	Name string
	age  int
}

//Sayok 外包可以访问
func (a *Acitons) Sayok() {
	fmt.Println("sayok", a.Name)
}

func (a *Acitons) hello() {
	fmt.Println("hello", a.Name)

}

//Bash 外包可以访问
type Bash struct {
	Acitons
}

func main() {
	var b Bash
	b.Acitons.Name = "jonny" // 或者简写 b.Name =   b.age =  因为b本身是个空的结构体，先找b没有属性name，age。就会直接去找actions的属性
	b.Acitons.age = 20
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	fmt.Println("b", b)
	log.Fatalln("这是一条会出发 fatal的日志")
	b.Sayok()
	b.hello()
	os.Exit(1)
	log.Panicln("这是一条会触发panic的日志。")
	// errors.New("error-New!")

}
