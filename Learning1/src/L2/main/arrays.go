package main
import "fmt"


// 定义方式：声明+类型
// var arr [3]int     // 声明初始化默认为 0   node :const array = []
// arr[0] = 1

// 声明同时赋值,一维数组   node ： new Array[1,2,3]  
// arra := [3]int{1,2,3}

// arrb := [3][2]int{{1,2,3},{4,5}}


func arrays(){

	var arr [3]int  // 3个长度的空数组
	arra := [3]int{1,2,3}  // 简写必须在方法内，
	arrb := [...]int{9,8,7,6,5,4,3}   // 类似于node 不定元素展开   const [...arryB] =arry
	arra[0]= 0001
	fmt.Println("arr",arr,"\narra:",arra,"\narrb:",arrb)
}


// 数组的遍历

func arrayfor(){
	arr := [...]int{99,98,96,7,6,5,4,3,2,1}
	// 常规用法 ,这里注意仍然是声明并赋值 := ,不是=，len方法== lenth
	for i := 0; i < len(arr); i++{
		fmt.Println("小标还是i",i,"\n对下表访问对应元素值",arr[i])
	}
}
// range 遍历
func  arrayrange(){
	arr := [...]int{100,200,300,400,500}
	// range 一个数组，分别声明并赋值给 两个变量， 如果不用某个变量 采用 _ 丢弃， for _, item := range arr{}
	for index, item := range arr{
		// 与node 数组遍历函数类似， index小标，item元素   Array.foreatch(item,idenx,arr){}
		fmt.Println("index",index,"\nitem",item)
	}
}


// 数据切片与大部分 用法一致， 但是不支持 负数 -1 -2




func main(){
	// arrays()
	// arrayfor()
	arrayrange()
}
