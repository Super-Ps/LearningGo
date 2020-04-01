// 结构体包裹的是字段定义
// 接口包括的是他的方法定义
// 接口类型声明中的这些方法就是改接口的方法集合
// 接口里的所有方法都没有方法体，接口的方法都是没有实现的方法，接口体现了程序设计的多态 和高内聚低耦合

//格式：
// type 接口类型  interface{

// 	方法名 1(参数列表)  返回值
// 	方法名 2(参数列表)  返回值
// }

// 几种接口无法实现的错误 ：
// 1. 函数名不一致
// 2. 实现接口的方法签名不一致  (签名部分： 方法名 1(参数列表)  返回值 ，与 method1(参数列表) 返回值列表{} 一致)
// 3. 接口中所有的方法均被实现，少1个都错误
// 什么是实现接口： 1. 把方法的内容实现出来  2.把所有接口的方法都实现了，
// fun (t 自定义类型) method1(参数列表) 返回值列表{}

package main

import (
	"fmt"
)

//DataWriter 定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct {
}

func (d *file) WriteData(data interface{}) error {

	fmt.Println("WriteData", data, "d", d)

	return nil
}

func main() {
	// 实例化file
	f := new(file)
	// 声明一个 DataWirter 的接口
	var w DataWriter

	// 讲接口赋值 f ，也就是 *file

	w = f

	// fmt.Println("W", w, "f", f)
	f.WriteData("test-----1")
	w.WriteData("test-data")
}
