package main

import (
	"fmt"
	"time"
)

func sendNumbers(ch chan int) {

	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)

}

func main() {

	ch := make(chan int, 1)

	go sendNumbers(ch)

	time.Sleep(1 * time.Second)
	fmt.Println(ch)
	// for v := range ch {
	// 	fmt.Printf("V=%d", v)
	// }
}
