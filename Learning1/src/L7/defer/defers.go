//todo 确保函数在结束时发生
//todo 参数在defer语句时计算
//todo  defer列表为后进先出
//todo  defer 讲语句放入栈时，也会将相应的值拷贝同时入栈，不会影响原有的变量值
package main

import (
	"bufio"
	"fmt"
	"os"
)

/**  块代码注释高亮用法
* *
* ?
* !

 */
func tryDefer() {

	defer fmt.Println("1") //? 即使有panic 也能执行
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	panic(" panic!") // * 输出： 3421  panic:  panic!
	fmt.Println("5")
}

//*参数在defer语句时计算
func tryDeferOld() {

	for i := 0; i < 100; i++ {

		defer fmt.Println(i)

		if i == 30 {
			panic("截止到")
		}
	}
}

func fibonace() func() int {

	x := 0
	y := 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func writeFile(filename string) {

	file, err := os.Create(filename) //! 创建文件

	if err != nil {
		panic(err)
	}
	defer file.Close() //! 关闭文件

	writer := bufio.NewWriter(file) //! 开辟一块内存，把数据缓存进去，当达到一定数量后一起写进去
	defer writer.Flush()            //! flush方法，就是强制刷新 缓冲区的一种方式。

	f := fibonace()

	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, f()) //! 格式化输
	}

}

func main() {
	// tryDefer()
	tryDeferOld()
	// writeFile("fibo.txt")
}

//todo 何时使用defer ?
//todo 在成对的操作时： open/close  Lock/Unlock  PrintHeader/PrintFooter
