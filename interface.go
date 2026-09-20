package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account int)
}

type Payment struct {
	// gateway stripe  // gateway instance is stripe
	gateway paymenter
}

func (p Payment) makePayment(amount float32) {
	p.gateway.pay(amount) //short of two lines
}

type razorPay struct{}

func (p razorPay) pay(amount float32) {
	//logic to make payment
	fmt.Println("Making payment using razorpay", amount)
}

type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account int) {
	fmt.Println("Making payment using paypal", amount)
	fmt.Println("Making payment using paypal", account)

}

func main() {
	// stripePaymentGw:=stripe{}
	// fakeGW:=fakePayment{}
	paypalGW := paypal{}
	newPayment := Payment{
		// gateway:stripePaymentGw,
		// gateway: fakeGW,
		gateway: paypalGW,
	}
	newPayment.makePayment(100)

}
