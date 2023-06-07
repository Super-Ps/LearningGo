package main

import (
	"fmt"
	"sync"
)

// 使用sync.Map避免并发访问map的问题：
// 查看bingfa6.go 的示例

func main() {

	var m sync.Map
	//存储键值
	m.Store("Key", "jonny")
	//加载（获取）值
	val, ok := m.Load("Key")
	if ok {
		fmt.Printf("Key,value: %s\n", val)
	} else {
		fmt.Printf("Key 不存在")
	}

	//删除键值对
	// m.Delete("Key")
	valB, ok := m.Load("Key")
	if !ok {
		fmt.Println("键值对被删除")
	}
	fmt.Println("value=", valB)

}
