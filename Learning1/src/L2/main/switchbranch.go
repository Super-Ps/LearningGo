package main
import "fmt"

// 与js不同的是需要显示break ,这里默认case 后带break，不需要显示写。
// 强制执行后面的代码就不再进行表达式的判断，直接执行
func switchBranch(i int){

	
	fmt.Println("int",i)
	switch {  // 这里可以不写表达式 之前写法 switch 表达式 {  }，现在直接 switch{}
	case i> 100:
		fmt.Println("i is equal to >100")
		 fallthrough  //  强制执行后面的代码 ，但是不能加在载最后一个分支
	case i< 100:
		fmt.Println("i is equal to <100, ")
		fallthrough
	case i<0:
		fmt.Println("i is equal to <0")
		fallthrough
	default:
		fmt.Println("All I know is that i is an integer")
		// fallthrough    比如这样		
	}
}

func main(){
	switchBranch(10)
}