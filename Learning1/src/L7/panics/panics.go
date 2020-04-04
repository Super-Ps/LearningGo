//todo panic:
//todo 停止当前函数执行
//todo 一直向上返回，执行每一层的defer
//todo 如果没有遇见recover 程序退出

//todo recover:
//todo 与panic 对应的 是recover
//todo 仅仅在defer 调用中使用
//todo 获取panic的值
//todo 如果无法处理 可重新处理panic

package main

import (
	"errors"
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()

		if err, ok := r.(error); ok {
			fmt.Println("ERROR ", err)
		} else {
			panic(r)
		}
	}()
	panic(errors.New("THIS,IS AN ERROR"))
}

func main() {
	tryRecover()
}
