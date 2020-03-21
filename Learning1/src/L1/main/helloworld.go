package main  	// 包， 表明代码所在的模块

import  "fmt" 
import   "os"   // 引入代码依赖

// 包内全局的声明,不同于js有全局变量

var (
	aa =0
	bb ="aaaaa"
)



func variable(){
// 定义变量  变量在前 类型在后
	var a int =3
	var b string ="abc"
	var c int
	var d string ="d"
	fmt.Println("a,b",a,b,c,d,cc,dd)

}
//  :=  形式定义变量  ,声明并赋值 ，如果只先声明需要var 
func variablenull(){

	a,b,c,d := 1,"bbb",true,1000000000

	b = "fffff"
	fmt.Println(":=",a,b,c,d)
}


//调用
func main()  {
	fmt.Println("命令行参数是 :",os.Args)
	fmt.Println(" Jony ")
	fmt.Println("Jony 222")
	variable()
	variablenull()
	os.Exit(0)
}


/*

 应用程序入口

	必须是 main 包    package main
	必须是main 方法    func main（）

	文件名不一定是 main.go
退出返回值
	Go中 不支持任何返回值
	通过os.Exit 返回状态

获取命令行参数
     mian 函数不支持参数
	os.Args 获取命令行参数 os.Args[0] 执行路劲为二进制文件   ，剩余后面就是参数了，


内建变量类型

*/





