package test_strcut

import (
	"fmt"
	"testing"
)

type person struct {
	name string
	age  int
}

func TestStrcut(t *testing.T) {

	p := person{"jonny", 30}

	fmt.Println("The person's name is ", p.name)
	// 按照顺序提供初始化值

}
