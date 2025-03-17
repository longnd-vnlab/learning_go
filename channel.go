package main

import (
	"fmt"
)

func main() {
	done := make(chan bool) // create a channel

	go func() {
		fmt.Println("running")

		done <- true // send signal done
	}()

	<-done
}
