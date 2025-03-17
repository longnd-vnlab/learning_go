package main

import (
	"fmt"
	"sync"
)

func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // Đánh dấu hoàn thành
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup

	messages := []string{"Hello", "Golang", "Concurrency"}
	for _, msg := range messages {
		wg.Add(1)
		go printMessage(msg, &wg)
	}

	wg.Wait() // Đợi tất cả goroutine hoàn thành
	fmt.Println("All goroutines finished")
}
