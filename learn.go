package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func main() {

	var a = 1
	var b = 2
	a, b = swap(a, b)
	fmt.Println(a, b)
}
