package main

import "fmt"

const (

	s1 = 60
	hour1 = s1 * 60
	day1 = hour1 * 24
)

// 返回多个值，支持对返回值命名   s hour day 是命名的返回值，命名之后retrun就不需要指定任何值返回，
// 遇到return时候直接返回命名
func resolveTime( input int) (s int, hour int,day int){
	
	s = input / s1
	hour = input / hour1
	day  = input / day1

	return 
}


// 匿名函数 跟js 一样
// 声明后调用
// func ( data int ){
// 	fmt.Println("hello",data)
// }(1000)
// 赋值给1个变量
//  f:= func ( data int ){
// 	 fmt.Println("hello",data)
//  }

 

//	 函数作为参数类型

type testFunc func(int )bool    // 声明一个函数类型


func isEven(intter int) bool{

	if intter % 2 ==0{
		return false
	}

	return true
}


func filter(slice []int, f testFunc) []int{
	result := []int{}
	for _, value := range slice {

		if f(value) {
			fmt.Println("if",f(value))
			result = append(result,value)
		}
	}
	return result
}







func main(){

	slice := []int{1,3,5,7,9,11,13,15,17,19}
	fmt.Println("sssss",)
	// resolveTime(1000)
	_,hour,day := resolveTime(180000)
	// f(200)
	fmt.Println("hour",hour,"day",day)

	// 函数作为参数传递

	even := filter(slice,isEven)
	fmt.Println("even",even)
}