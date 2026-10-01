package main

import "fmt"

type Payment struct {
	gateway  stripe
	gateways RazorPay
}

func (p Payment) MakePayment(amount float32) {
	// to get interferences of RazorPay struct

	// razorpayPaymentGW := RazorPay{}
	// razorpayPaymentGW.Pay(amount)

	// stripePayment := stripe{}
	// stripePayment.Pay(amount)
	// p.gateway.Pay(amount)

	p.gateway.Pay(amount)
	p.gateways.Pay(amount)

}

type RazorPay struct{}

func (r RazorPay) Pay(amount float32) {
	// logic to make Payment
	fmt.Println("making Payment using RazorPay", amount)
}

type stripe struct{}

func (s stripe) Pay(amount float32) {
	fmt.Println("making Payment using stripe", amount)
}

func main() {
	// used by stripe payment
	stripePaymentGw := stripe{}
	newPayment := Payment{
		gateway: stripePaymentGw,
	}
	newPayment.MakePayment(100)

	// to use the RazorPay
	razorpayPaymentGW := RazorPay{}
	newPayments := Payment{
		gateways: razorpayPaymentGW,
	}
	newPayments.MakePayment(200)
}
