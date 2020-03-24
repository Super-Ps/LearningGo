package main
import ("fmt"
		// "strconv"
		)

//  声明格式 map[KeyType]ValueType
// m := map[string]int{"one":1,"twO":2}
// m := map[string]int{}
// m1["one"] = 1
// make(map[string]int)

 func mapinit(){

	m1 := map[string]int{"a":1,"b":2,"c":3}
	m2 := map[int]int{10:100,20:200,30:300}
	m3 := map[string]int{}

	m4 := make(map[string]int,5)

	fmt.Println("m1",m1["a"],"\nm2",m2[30],"\nm3",len(m3),"\nm4",len(m4)) //
 }


//  在访问的key不存在时，仍然会返回0值，不能通过返回 nil 来判断元素是否存在
 func  mapok(){ 

	m := make(map[string]int,3)
	n := make(map[int]int,4)
	fmt.Println("m",m) 
	fmt.Println("n",n) 
	fmt.Println("n",n[1])  // 0  不存在的key赋值为 0

	
	 n[2] = 0
	fmt.Println("n",n,len(n))  //  map[2:0]     1
	m["a"] = 10
	v ,ok  := m["a"];   // map 返回值 返回 key值，跟是否存在的布尔值，对比别的语言不存在就报错，

	fmt.Println("out-v",v,"out-ok",ok) //  out-v 10 out-ok true

	if v ,ok := m["a"]; ok {      // 这里可以声明一个变量
		fmt.Println("v",v,"ok",ok)

	}else {
		fmt.Println("error-1")
	}

	if v ,ok := n[1]; ok {
		fmt.Println("v2",v,"ok2",ok)
		
	}else {
		fmt.Println("error-2")
		fmt.Println("v2",v,"ok2",ok) // 不存在复制为0  布尔值为false   v2 = 0 ok2  =false
	}
 }

// range  遍历map  k ,v   ,map 的遍历是无序的。每次运行的顺序都不一样
 func maprang(){

	m := map[string]int{"a":1,"b":2,"c":3}

	for k ,v := range m {
		fmt.Println("key",k,"value",v)
	}
 }
  


 //  Map 的值 除了是数据类型 也可以是一个方法
  
 func mapvalueIsFuc(){

	m := map[int]func( input int)int{}
	m[1] = func(input int) int{ return input}
	m[2] = func(input int) int{ return input*input}
	m[3] = func(input int) int{ return input*input*input}

	n := m[1]
	fmt.Println("m1",m[1](2),"\nm2",m[2](2),"\nm3",m[3](2),"\nn",n)
 }


 // go 没有set类型，通过map 实现 set, 
  func mapSet(x int){

	mySet := map[int]bool{}
	
	mySet[1] = true
	mySet[2] =false

	n := x
	if mySet[n] {
		fmt.Println("%d is existing",n)	
	}else{
		fmt.Println("%d is existing",n)
	}

	fmt.Println("长度",len(mySet))

	

  }
 func main(){
	// mapinit()
	// mapok()
	// maprang()
	// mapvalueIsFuc()
	mapSet(2)
 }