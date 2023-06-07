package main

import "fmt"

func channelTest() {
	ch1 := make(chan int, 1)
	ch1 <- 2

	select {
	case v := <-ch1:
		fmt.Println("取到的数据", v)
	case ch1 <- 1:
		fmt.Println("写入数据")
	}

}

// 从关闭的且没有值的channel 中取值
func closechannelTest() {

	//声明实例化通话ch1
	ch1 := make(chan int, 1)
	//关闭通道
	close(ch1)

	select {
	case v := <-ch1:
		fmt.Printf("closechannelTest ---从被关闭的ch1中取值:%d \n", v)
	default:
		fmt.Println("默认case")
	}
}

// 从关闭的且有值的channel 中取值
func closechannelvalueTest() {

	//声明实例化通话ch1
	ch1 := make(chan int, 1)
	//向通道中赋值
	ch1 <- 20230607
	//关闭通道
	close(ch1)
	// 关闭之后取值
	v1 := <-ch1
	fmt.Printf("closechannelvalueTest---关闭之后取值的: %d\n", v1)
	select {
	case v := <-ch1:
		fmt.Printf("从被关闭的ch1中取值:%d", v)
	default:
		fmt.Println("默认case")
	}
}
func main() {
	channelTest()
	closechannelTest()
	closechannelvalueTest()
}
