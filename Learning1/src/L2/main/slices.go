package main

import "fmt"

/*
声明方式1：
slice 与数组 的申请方式相似， 数组是必须给定数量: a := [...]int{1,2,3}。而slice是动态的可以为空 a := []int{1,2,3}

声明方式2：
make(slice,len,cap)

len() 和 cap() 函数
切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
*/

func slices() {

	// 声明一个slice  var  a []T
	var slicea []string

	slicea = append(slicea, "go") // 往原切片的末尾添加元素。类似 node push   python append
	slicea = append(slicea, "node")

	// 声明初始化一个slice

	s2 := []int{1, 2, 3, 4, 5}

	// make初始化一个slice
	s3 := make([]int, 5, 7)

	s3[4] = 100 // 给某个位置赋值，不能超过最大长度。

	fmt.Println("slicea", slicea, len(slicea), cap(slicea))
	fmt.Println("s2", s2, len(s2), cap(s2))
	fmt.Println("s3", s3, len(s3), cap(s3))
	fmt.Println("查看长度和容量", s3[0], s3[1], s3[2], s3[3], s3[4]) //
}

// 切片是一个共享数据结构
func slicePublick() {
	zimu := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	news := zimu[2:4] //下标2开始 到3结束，不包含4

	news2 := zimu[4:7]
	fmt.Println("zimu", zimu, len(zimu), cap(zimu))
	fmt.Println("news", news, len(news), cap(news)) //[C D] 2 7   这里的cap容量7 表示从第2位开始到最后一位的总长度
	fmt.Println("news2", news2, len(news2), cap(news2))
	news[0] = "xxx"
	fmt.Println("news", news, len(news), cap(news))
	fmt.Println("news2", news2, len(news2), cap(news2))
	fmt.Println("zimu", zimu, len(zimu), cap(zimu)) // zimu [A B xxx D E F G H I] 9 9  新切片数组被赋值后会影响共享的slice，

	news3 := zimu[3:6]
	fmt.Println("news3", news3, len(news3), cap(news3))
	news3[0] = "YYY"
	fmt.Println("news3", news3, len(news3), cap(news3)) // [YYY E F] 3 6
	fmt.Println("zimu", zimu, len(zimu), cap(zimu))     // [A B xxx YYY E F G H I] 9 9

}

// 数组 可比较，切片不能

func arryaSlice() {
	// 数组
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}

	if a == b {
		fmt.Println("数组相等", a == b)
	}

	// 切片  ，不能比较 会报错
	c := []int{1, 2, 3}
	d := []int{1, 2, 3}
	if c == d {
		fmt.Println("切片不能比较", c == d) //invalid operation: c == d (slice can only be compared to nil
	}
}

func main() {
	// slices()
	// slicePublick()
	arryaSlice()
}
