//select 多路复用

package main

import (
	"fmt"
	"time"
)

func sendNumbers(ch1 chan int) {
	for i := 1; i <= 5; i++ {
		ch1 <- i
	}
	close(ch1)
}

func sendLetters(ch2 chan string) {
	for i := 'a'; i <= 'e'; i++ {
		ch2 <- string(i)
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 3)
	ch2 := make(chan string, 3)

	go sendNumbers(ch1)
	go sendLetters(ch2)

	for {
		select {
		case n, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				fmt.Println("数字:", n)
			}
		case c, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				fmt.Println("字母:", c)
			}
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}

	time.Sleep(1 * time.Second)
}
