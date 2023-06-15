package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 复合类型（数组类型）反射
	arr := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]string{"A", "a", "b", "c"}
	typ := reflect.TypeOf(arr)
	fmt.Println("kind:", typ.Kind()) //kind: array
	typ2 := reflect.TypeOf(arr2)
	fmt.Println("kind:", typ2.Kind()) //kind: array

	//指针类型
	var p1 *float64 // 申请的一个指针类型
	var p2 *string
	ptyp := reflect.TypeOf(p1)
	fmt.Println("kind:", ptyp.Kind()) //kind: ptr
	ptyp2 := reflect.TypeOf(p2)
	fmt.Println("kind:", ptyp2.Kind()) //kind: ptr

	zhengshu := 10
	zhizhenbianliang := &zhengshu //取变量的地址
	fmt.Println(" &zhengshu", &zhengshu, zhizhenbianliang)
	typp := reflect.TypeOf(zhizhenbianliang)
	fmt.Println("变量的地址类型是", typp.Kind()) // 指针地址的类型是 ptr
	fmt.Println("变量本身类型", typp.Elem())   //   变量本身类型 int

	// 切片类型
	s := make([]int, 5, 10)
	typs := reflect.TypeOf(s)
	fmt.Println("kind:", typs.Kind()) //kind: slice
	//typs.Len()  用类型调用长度和容量会报错，说明长度和容量并不带切片类型范畴，
	//而是与切片变量值绑定的，下面的示例印证了这一点：

	vas := reflect.ValueOf(s)
	fmt.Println("valueof:", vas.Len()) //valueof: 5
	fmt.Println("valueof:", vas.Cap()) // valueof: 10

	// map类型 ，它是一个无序键值对的集合，通过反射可以获取其键和值的类型信息：
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	typm := reflect.TypeOf(m)
	fmt.Println("kind:", typm.Kind())                 // 类型是 kind: map
	fmt.Println("key:", typm.Key())                   // key的类型是 key: string
	fmt.Println("Elem:", typm.Elem())                 // value的类型是 Elem: int
	fmt.Println("获取map的长度", reflect.ValueOf(m).Len()) // 拿到这个map的长度，直接用valueof，也是同切片一样

}
