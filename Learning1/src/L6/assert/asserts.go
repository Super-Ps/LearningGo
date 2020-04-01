package main

import "fmt"

//引出类型断言的 例子
//Ponit int
type Ponit struct {
	x int
	y int
}

func main() {
	var a interface{}
	var p Ponit = Ponit{1, 2}
	a = p  // 空接口可以 接收任意类型
	fmt.Println("p=", p, "a=", a)
	var p1 Ponit
	// p1 = a  // ERROR !  cannot use a (type interface {})  as type Ponit in assignment: need type assertion
	// p1 =a 不能直接复制，需要一个类型断言来解决
	p1 = a.(Ponit)  // 表示判断a 是否指向Point 类型的变量，如果是就转成Ponit 类型，并复制给b变量，否则报错
	fmt.Println("p1", p1)

	// 类型断言，用检测
	if x,ok := a.(Ponit);ok{
		fmt.Println("转成功",x,ok)
	}else{
		fmt.Println("转失败")
	}
}
// .(类型) 类型判断可以对已有的类型判断，也可以对自定义类型判断
//进行类型断言时， 如果不匹配，就会包panic，因此进行类型断言时，要确保原来的空接口指向的就是断言的类型
