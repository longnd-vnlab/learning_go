package main

import (
	"fmt"
	"time"
)

// Hàm đơn giản in ra 5 lần
func printMessage(message string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(message, i)
		time.Sleep(time.Millisecond * 500) // Giả lập công việc mất thời gian
	}
}

func main() {
	// Chạy đồng thời hai hàm bằng Goroutines
	go printMessage("Hello from Goroutine 1")
	go printMessage("Hello from Goroutine 2")

	// Đợi một lúc để Goroutines chạy (thực tế nên dùng sync.WaitGroup)
	time.Sleep(time.Second * 3)
	fmt.Println("Main function finished")
}
