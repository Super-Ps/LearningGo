
//  string 是数据类型不是引用或者指针 
//  string 是只读的byte slice，（string 是一个不可变的 ） len函数 打印出来的是byte的个数
//  string 的byte数组可以放任何数据 

package main

import("fmt"
		"strings"
		"strconv"
	)



func stringTest(){

	var s string
	t := "我是字符串"  // 一个中文字符占3个byte   

	s ="我有不同的情况：:!@#AbcabC"
	fmt.Println("%d t的长度",len(t),"%d s的长度",len(s))  // 15  , 34

	s = "hello"  
	fmt.Println("%d t的长度",len(s))  // 5  一个字母 一个byte
	s ="\xE4\xBB\xA5"  // 二进制数据

	fmt.Println("%d t的长度",len(s))  // 二进制有3个部分组成的 长度 是 3 

	// s[1] = "xA5"    	//不能修改 所以不可变
	// fmt.Println("%d s",s)   //   cannot assign to s[1]	
}



func   runneTest(){

	r :="中文"
	fmt.Println("%d r",len(r),"%x ",r)  //  长度6   
	s :=[]rune(r)  				// 通过rune 取出字符串里的unicode 
	fmt.Println("%d s",len(s),s)   // 长度2  

	zh :="放弃python，转向Go"
	
	for _ ,v := range zh{
		fmt.Println("%d %T\n ",v)
	}
}


// 常用的字符串包
 func  stringPackage(){
	s :="世界你好,Q,W,E,R"
	news := strings.Split(s,",")   // 字符分割 首写大写
	news2 := strings.Join(news,"-")	  // 连接
	fmt.Println("news",news)  // 一个数组
	fmt.Println("news",news2)  // 一个数组
	for _ ,item := range news{
		fmt.Println("item",item)
	}
 }


 // 常用的类型转换
 func strconvTest(){

	s := strconv.Itoa(10000)  // int 转字符

	fmt.Println("s",s)
	
	// i ,err:= strconv.Atoi("9999S")   // 返回2个值
	// fmt.Println("i",i, err)

	if i ,err:= strconv.Atoi("9999S");err == nil {
	fmt.Println( i +1 )
	}else{
		fmt.Println("转换失败")
	}
 }

func  main(){
	// stringTest()
	// runneTest()
	// stringPackage()
	strconvTest()

}