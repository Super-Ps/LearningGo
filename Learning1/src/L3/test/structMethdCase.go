package main

import "fmt"

//CASE1
type Person struct {
	Name string
}

// CASE2
type Circle struct {
	redius float64
}

func (person Person) spkeak() {
	fmt.Println(person.Name, "是一个好人")
}

func (person Person) jisuan() {
	res := 1
	person.Name = "miyas"
	for i := 0; i <= 1000; i++ {

		res += i
	}
	fmt.Println(person.Name, "计算结果是：", res)
}

func (person Person) jisuan2(n int) {

	res := 1
	person.Name = "miyas"
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println(person.Name, "计算结果是：", res)
}

func (person Person) getSum(n1 int, n2 int) int {

	return n1 + n2
}

func (c Circle) area() float64 {

	return 3.14 * c.redius * c.redius
}

// 方法通常和结构体的指针类型绑定

func (cc *Circle) area2() float64 {
	// 因为	CC 是指针，因为我们标准的访问其字段的方式是 (*c).redius ,简写为  c.redius

	fmt.Println("cc 此时已经是指针，他指向的地址是*Circle类型的地址 ", cc) // 不能用取地符取 &，因为在外面传进来已经是指针了

	fmt.Printf("C 是 *Circle 指向的地址 =%p", cc)
	cc.redius = 11.000 // 内部修改1次
	return 3.14 * cc.redius * cc.redius
}

func main() {
	p := new(Person)

	p.Name = "Jonny"

	p.spkeak() // 通过一个变量去调用方法时，其调用机制和函数一样
	p.jisuan() // 不一样的地方时，变量调用方法时，改变量本身也会作为一个参数传递到方法，（如果 变量是值类型，则进行值拷贝
	// 如果变量是 引用类型，则进行地址 拷贝）
	p.jisuan2(100)
	res := p.getSum(100, 200)
	fmt.Println("case1", res)

	c := new(Circle)
	c.redius = 5.0
	res2 := c.area()
	fmt.Println("case2", res2)

	// 指针传递
	cc := new(Circle) // 返回的是一个指针 cc
	// fmt.Println("cc的地址",&cc)
	fmt.Printf("cc 的地址 =%p ", cc) // 0xc00001a0b0 与上面传进去的 cc 变量指向的地址 是 0xc00001a0b0 ，说明是操作同一块内存
	cc.redius = 22.0              // 赋值1次
	res3 := cc.area2()
	fmt.Println("case3", res3, cc.redius)


}
