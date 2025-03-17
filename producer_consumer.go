package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Produced:", i)
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(ch) // Đóng channel sau khi gửi xong
}

func consumer(ch chan int) {
	for value := range ch {
		fmt.Println("Consumed:", value)
	}
}

func main() {
	ch := make(chan int, 3) // Channel có buffer = 3

	go producer(ch)
	go consumer(ch)

	time.Sleep(time.Second * 3) // Chờ đủ thời gian để xử lý
}
