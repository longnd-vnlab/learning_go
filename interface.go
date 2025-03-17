// Giả sử ta có hệ thống quản lý thanh toán, có nhiều phương thức thanh toán khác nhau (thẻ tín dụng, PayPal, v.v.):

package main

import "fmt"

// interface payment
type Payment interface {
	pay(amount float64)
}

// struct creditcard
type creditcard struct {
	number string
}

// function pay
func (c creditcard) pay(amount float64) {
	fmt.Println("Paying", amount, "with credit card", c.number)
}

// struct Paypal
type Paypal struct {
	email string
}

// function pay for credit card
func (p Paypal) pay(amount float64) {
	fmt.Println("Paying", amount, "with Paypal", p.email)
}

func paymentProcess(p Payment, amount float64) {
	p.pay(amount)
}
func main() {
	cc := creditcard{
		number: "1234567890",
	}

	cc.pay(200)
}
