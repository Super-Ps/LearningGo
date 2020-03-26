package main

import "fmt"

//
type People struct {
	name string

	child *People // 类型是 *People
}

//键值对初始化结构体的书写格式 K : V
var relations = &People{
	name: "Jonny",
	child: &People{
		name: "miya",
		child: &People{
			name: "月月",
		},
	},
}

func main() {
	fmt.Println("relations", relations)
}
