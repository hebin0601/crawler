package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("Hello from "+"goroutine %d\n", i)
		}(i)
	}
	time.Sleep(10 * time.Millisecond)
}
