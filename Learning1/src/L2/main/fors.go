package main


import "fmt"

// 与node区别 去处了变量申请关键字，和（） for(let i =0;i<10;i++)
func fors(){
	sum := 0
	for i:= 0; i<10; i++{
		sum += i
		fmt.Println("i",i)
	}
	fmt.Println("sum",sum)
}

// 省略表达式1，和表达式3 就变成了 while循环

func while(){
	sum := 0
	// while(){}
	for sum<10{
		sum += sum
		fmt.Println("sum",sum)
		 sum++
	}
}
func main(){
	fors()
	while()
}