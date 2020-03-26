package main

import "fmt"


type Person struct{

	Name string
}


func ( person Person) spkeak(){
	fmt.Println(person.Name,"是一个好人")
}

func (person Person)jisuan(){
	res :=1
	person.Name ="miyas"
	for i:= 0;i<=1000; i++ {

		res += i
	}
	fmt.Println(person.Name,"计算结果是：",res)
}

func (person Person) jisuan2(n int){

	res :=1
	person.Name ="miyas"
	for i:= 0;i<=n; i++ {
		res += i
	}
	fmt.Println(person.Name,"计算结果是：",res)
}

func (person Person) getSum(n1 int,n2 int) int{

	return n1 + n2
}



func  main()  {
	p := new(Person)

	p.Name = "Jonny"

	p.spkeak()
	p.jisuan()
	p.jisuan2(100)
	res := p.getSum(100,200)
	fmt.Println("res",res)
}