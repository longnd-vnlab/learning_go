package main

import (
	"fmt"
	"sync"
)

var count int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	count++ // Không có bảo vệ, có thể xảy ra lỗi race condition
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go increment(&wg)
	go increment(&wg)

	wg.Wait()
	fmt.Println(count) // Kết quả không đảm bảo chính xác
}
