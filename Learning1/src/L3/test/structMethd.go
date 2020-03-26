// 结构体除了字段还需要有行为，所以需要用到方法，与函数很相似
// go 中方法的作用在指定的数据类型上的 (和指定的数据类型绑定)，因此自定义类型，都可以有方法，而不仅仅是struct
package main
import "fmt"

type Person struct{  

	Name string
	age int
}
// 这里的p 表示 哪个 person 变量调用，这个p就是调用者的副本
func (pp Person) test(){      // A 结构体有一个方法，方法名叫test   (a A) 体现test方法和A 类型绑定

	fmt.Println("func pp",pp,"地质",&pp)
	pp.Name= "kobi"
	 fmt.Println( "in pp",pp,pp.Name,"地质",&pp)
	
}

func  main()  {
	
	// var p Person
	pp := new(Person)  
	fmt.Println("实例化后的pp",pp,&pp)
	pp.Name = "Jonny"
	pp.test()    // 调用时 相当于这个p 传递进入方法内的 类型 p 
	// test方法只能与 Person类型的变量来调用，不能直接调用 test(),也不能给别的 类型来调用
	fmt.Println("out pp",pp,"地质",pp.Name,&pp)
	
}