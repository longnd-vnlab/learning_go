package main

import "fmt"

var a = []int{1, 2, 3}

var b = [10]int{
	1, 5: 99, 9: 100,
} // this called sparse array, the first element is 1, the fifth element is 99, the ninth element is 100, the rest of the elements are 0

// Khi biến array được gán hoặc truyền, thì toàn bộ array sẽ được sao chép. Nếu kích thước của array lớn, thì phép gán array sẽ chịu tổn phí lớn. Để tránh việc overhead (tổn phí) trong việc sao chép array, bạn có thể truyền con trỏ của array.
var x = [...]int{1, 2, 3}

var y = &x

func main() {
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(y[0], y[1])

	for i := range y {
		fmt.Println(i)
	}
}
