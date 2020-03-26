package main

import (
	"fmt"
)

// 结构体默认为值 拷贝，值类型
//  如果结构体的类型是 指针 * ,slice ,map的零值都是nil， 即还没有分配空间
type person struct {
	Name   string
	Age    int
	Scores [5]float64
	Ptr    *int
	Slice  []int             // 切片必须先make
	Maps   map[string]string // 切片必须先make
}

type personB struct {
	Name string
	Age  int
}

func main() {
	// 赋值给不同类型 1.直接声明   var 变量名 结构体名
	var p person
	fmt.Println("The person's name is ", p) // The person's name is  { 0 [0 0 0 0 0] <nil> [] map[]}

	p.Slice = make([]int, 4)
	p.Slice[0] = 100
	fmt.Println("The person's name is ", p) // The person's name is  { 0 [0 0 0 0 0] <nil> [100 0 0 0] map[]}

	p.Maps = make(map[string]string)
	p.Maps["key"] = "mapValue"
	fmt.Println("The person's name is ", p) //The person's name is  { 0 [0 0 0 0 0] <nil> [100 0 0 0] map[key:mapValue]}

	// 方式2

	p1 := personB{}
	p1.Name = "kobi"
	fmt.Println("The person's name is ", p1)
	p2 := personB{"kobi2", 20}
	fmt.Println("The person's name is ", p2)
	// 方式3

	p3 := new(personB) //   创建指针类型的结构体 也叫 创建实例访问属性 ，p3 = *p3,操作方便简写为p3

	p3.Name = "Aonier"

	fmt.Println("The person's name is ", p3) //  &{Aonier 0}

	// 方式4  取结构体的地址 实例化

	p4 := &personB{}
	p4.Name = "取地址符 &"

	fmt.Println("The person's name is ", &p4)
}
