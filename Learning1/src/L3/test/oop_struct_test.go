package test_strcut

import ("testing"
		"fmt"
)


type person struct {

	name string
	age int
}

var P person   // P1现在就是person类型的变量
	P.name = "Jonny"
	P.age = "Jonny"



func  TestStrcut(t *testing.T){

	fmt.Println("The person's name is %s", P.name)
	// 按照顺序提供初始化值
	
	
}