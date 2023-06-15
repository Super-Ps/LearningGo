package main

// channel 反射的用法
import (
	"fmt"
	"reflect"
)

func main() {
	//channel是Go特有的类型，channel与切片很像，
	//它的类型信息包括元素类型、chan读写特性，但channel的长度与容量与channel变量是绑定的，看下面示例：
	ch := make(chan<- int, 10)
	ch <- 1
	ch <- 2
	typch := reflect.TypeOf(ch)
	fmt.Println("初始化的这个ch类型---kind", typch.Kind())     // chan
	fmt.Println("channel内写入的变量类型--Elem", typch.Elem()) // int
	fmt.Println("ChanDir", typch.ChanDir())

	// 高级用法：https://mp.weixin.qq.com/s/E4lT4SuWKIlCZd60i7vigQ 待实践
}
