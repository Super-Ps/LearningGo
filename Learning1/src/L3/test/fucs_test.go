package test_fuc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 结构示例

/*
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
    //这里是处理逻辑代码
    //返回多个值
    return value1, value2
}
*/

// 值传递，传递的是 X 本身的copy，改变的是 这个x的copy值，不会改变 X 本身
func max(a, b int) int {

	if a > b {
		a = a + b
		return a
	}
	return b
}

// 传指针
func maxZhiZhen(a *int) int {
	*a = *a + 1 // 修改了a的值
	return *a
}

// 返回多个值

func returnMoreValue() (int, int) {
	return rand.Intn(100), rand.Intn(200)
}

// 输入一个函数类型 ，返回一个函数类型 .类似于装饰器，给一个函数包装一个功能

func timeCall(inner func(opt int) int) func(opt int) int {

	return func(n int) int {
		start := time.Now()

		fmt.Println("start", start)
		res := inner(n)
		fmt.Println("res", res, "N", n)
		fmt.Println("time ", time.Since(start).Seconds())

		return res
	}
}

func slowFunc(op int) int {
	fmt.Println("time.Second", time.Second)
	// time.Sleep(time.Second * 1)
	return op
}

// 可边长参数 （变量名  ...T）
func sumTest(in ...int) int {

	sum := 0
	for _, itme := range in {
		sum += itme
	}
	return sum
}


// 将传入的秒解析为3种时间单位


// 函数作为值，类型

func TestFuc(t *testing.T) {
	// a, b := returnMoreValue()
	// t.Log(a, b)

	// timeR := timeCall(slowFunc)
	// t.Log(timeR(100))

	t.Log("sumTest=", sumTest(100, 200, 300, 4000))

	x := 10
	y := 5

	zhizhen := 100
	//	 值传递
	max_return := max(x, y)
	t.Log("max_return", max_return, x, y)

	// 传递 变量地址
	newzhizhen := maxZhiZhen(&zhizhen) //调用 add1(&x) 传x的地址
	t.Log("newzhizhen", newzhizhen)
	t.Log("zhizhen", zhizhen)

}
