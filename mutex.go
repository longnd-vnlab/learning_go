package main

import (
	"fmt"
	"sync"
)

var (
	count int
	mu    sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	count++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go increment(&wg)
	go increment(&wg)

	wg.Wait()
	fmt.Println(count) // Kết quả luôn chính xác
}
