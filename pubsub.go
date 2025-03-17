package main

import (
	"fmt"
	"sync"
)

type PubSub struct {
	subscribers []chan string
	mu          sync.Mutex
}

func (ps *PubSub) Subscribe() chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string, 1)
	ps.subscribers = append(ps.subscribers, ch)
	return ch
}

func (ps *PubSub) Publish(message string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for _, ch := range ps.subscribers {
		ch <- message
	}
}

func main() {
	ps := &PubSub{}

	sub1 := ps.Subscribe()
	sub2 := ps.Subscribe()

	go func() {
		fmt.Println("Subscriber 1 received:", <-sub1)
	}()

	go func() {
		fmt.Println("Subscriber 2 received:", <-sub2)
	}()

	ps.Publish("Hello, Subscribers!")
}
