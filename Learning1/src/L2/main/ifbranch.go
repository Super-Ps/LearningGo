package main
import "fmt"
import "io/ioutil"  // 读文件库



// IF 与其他语言比较 不表达式没有()

// if x > 10 {
//     fmt.Println("x is greater than 10")
// } else {
//     fmt.Println("x is less than 10")
// }

// IF  嵌套 示例
func branchIf(){
	
	a ,b := 100 ,200
	
	if a ==100{
		fmt.Println("执行第一个if")
		if b == 200{
			fmt.Println("执行第二个if")
		}
	}

	fmt.Println("a的值:%d",a,"\nb的值：%d",b)
}


// 读文件示例
// nil 的理解：不等同于node 里的null , 这里只是 指针、切片、映射、通道、函数和接口的零值则是 nil，
//其他基础类型(对比js 来理解):比如布尔初始值是flase,数值类型初始值0,字符串初始值为"""
// 先这样理解老熟悉if结构 后面再研究，
func  readIo(){

	const filename = "abc.text"
	contens,err :=ioutil.ReadFile(filename)
	if err != nil{       
		fmt.Println(err)
	}else{
		fmt.Printf("%s普通写法:\n",contens) // 小写s打印字符串，不是大写
	}

}



// Go的if还有一个强大的地方就是条件判断语句里面允许声明一个变量，
//这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了

func readIoTwo(){
	const filename = "abc.text"

	if contens,err :=ioutil.ReadFile(filename);err !=nil{
		fmt.Println(err)

	}else{
		fmt.Printf("%s声明写法",contens)
	}
}


//相当于js,把上面的条件值，拿到表达式里面写，后面加上;再写条件
// if(let a=0,a<b){

// }

func main(){
	branchIf()
	readIo()
	readIoTwo()
}