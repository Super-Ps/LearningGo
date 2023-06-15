package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	Speak() string
}

type Cat struct{}

func (c Cat) Speak() string {

	return "Meow"
}
func (c Cat) Speaka() []int {

	return []int{1, 2, 3, 4, 5}
}

func main() {
	var a Animal = Cat{}
	typ := reflect.TypeOf(a)
	fmt.Println(typ.Kind())         // interface
	fmt.Println(typ.NumMethod())    // 1
	fmt.Println(typ.Method(1).Name) // Speak ,Speaka  Method(0)从索引内指定要获取哪个方法名字
	fmt.Println(typ.Method(0).Type) // func(main.Animal) string
}
