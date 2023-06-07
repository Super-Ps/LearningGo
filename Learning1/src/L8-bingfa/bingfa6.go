package main

// sync.Map是为了解决在并发场景下普通map非线程安全的问题而引入的。
// 普通的Go map在多个goroutine同时读写时可能会出现竞争问题，这时候需要使用互斥锁或者读写锁来保护map。
// 而sync.Map内部实现了这种保护机制，在并发读写时确保线程安全，让开发者能更方便地使用map来存储数据。

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	var wg sync.WaitGroup
	totalWorkers := 8

	// 写入数据的worker协程
	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sm.Store(fmt.Sprintf("worker-%d", id), id)
			fmt.Printf("Worker-%d 写入数据\n", id)
		}(i)
	}

	// 读取数据的worker协程
	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("worker-%d", id)
			value, ok := sm.Load(key)
			if ok {
				fmt.Printf("读取数据: %v => %v\n", key, value)
			}
		}(i)
	}

	wg.Wait()
}
