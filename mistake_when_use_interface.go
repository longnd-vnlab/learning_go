package main

import "fmt"

// Định nghĩa interface lỗi runtime
type runtimeError interface {
	error
	RuntimeError()
}

// Một struct có thể tự động implement runtimeError nếu nó có method `RuntimeError`
type MyError struct{}

func (e MyError) Error() string {
	return "Lỗi!"
}

func (e MyError) RuntimeError() {}

func main() {
	var err runtimeError = MyError{}
	fmt.Println(err.Error()) // Lỗi!
}
