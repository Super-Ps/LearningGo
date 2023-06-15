package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 基本类型反射
	s := "hello,world"
	i := 2880
	t := reflect.TypeOf(s) //TypeOf函数获取一个字符串的类型信息
	fmt.Println(t.Name())
	ti := reflect.TypeOf(i)
	fmt.Println(ti.Name())

	vi := reflect.ValueOf(i)
	fmt.Println("通过值的类型方法来获取对应的值，这里是int类型", vi.Int())
	vs := reflect.ValueOf(s)
	fmt.Println("通过值的类型方法来获取对应的值，这里是string类型", vs.String())
}
