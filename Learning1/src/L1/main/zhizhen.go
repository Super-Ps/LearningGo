package main

import (
	"fmt"
)

/*
整形
fmt.Printf("%b\n", n)
fmt.Printf("%c\n", n)
fmt.Printf("%d\n", n)
fmt.Printf("%o\n", n)
fmt.Printf("%x\n", n)
fmt.Printf("%X\n", n)




字符串  []byte

s := "小王子"
fmt.Printf("%s\n", s) // 直接输出字符串 或者 []byte   小王子
fmt.Printf("%q\n", s) // 该值对应的双引号括起来的go语法字符串字面值  "小王子"
fmt.Printf("%x\n", s) // 16进制  a-f
fmt.Printf("%X\n", s) // 16进制  A-F



指针

a := 10
fmt.Printf("%p\n", &a)   // 标识为16进制，并加上前导 0X   0xc000094000
fmt.Printf("%p\n", &a)
*/

func f() *int {

	v := 1
	return &v
}

func incr(p *int) int {
	fmt.Println(" p的值", p) // p的值 0xc00001a0d0
	*p++                   // 改变值
	fmt.Println(" *p的值", *p)
	return *p
}

//! better comments
//? 蓝色
//// 删除不用
//todo 橘色
//* 绿色
func main() {

	var x int = 10
	a := &x
	fmt.Println("\n", &x)
	fmt.Println("a", *a)
	fmt.Println("指针类型的零值是 nil", a == nil, "指针的值可以比较的，指向同一个变量或者都是nil时才相等", &a == &a, &a == nil)
	//
	fmt.Println("每次调用都返回一个不同的值", f() == f())
	//
	v := 100
	fmt.Println(&v) // 0xc00001a0d0
	incr(&v)
	//

}
