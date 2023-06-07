package main

import (
	"context"
	"fmt"
	"time"
)

// 使用Go context来实现goroutine的取消和超时：

func worker(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("工作退出：", ctx.Err())
	case <-time.After(3 * time.Second):
		fmt.Println("工作完成")
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go worker(ctx)

	time.Sleep(8 * time.Second)

	fmt.Println("程序结束")
}
