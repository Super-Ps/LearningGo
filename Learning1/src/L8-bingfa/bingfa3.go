// 使用sync.waitgroup 同步多个goroutines

// 在Go中，sync.WaitGroup主要用于解决goroutine同步的问题。当你使用并发执行多个任务时，经常需要等待所有协程都完成后再继续执行后续工作（比如汇总结果、关闭资源等）。这时候就可以使用sync.WaitGroup来确保所有协程都完成任务后，才执行主goroutine剩下的操作。

// sync.WaitGroup具有以下特点：

// 具有一个计数器（初始值为0），用于记录未完成任务的协程数。
// 提供了Add(delta int)函数来增加待完成任务计数。
// 提供了Done()函数来将待完成任务计数减1，通常通过defer调用确保计数器在协程结束时减1。
// 提供了Wait()函数用于阻塞直到计数器值为0，表示所有任务都已完成。
package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, id int) {
	fmt.Printf("this is %d worker\n", id)
	defer wg.Done()
}

func main() {
	numvbes := 20
	var wg sync.WaitGroup

	for i := 0; i < numvbes; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}

	wg.Wait()
	fmt.Println("所有任务已完成")
}
