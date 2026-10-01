package main

import "fmt"

type payment struct {
	// gateway stripe
	gateway paymenter
}
// Open close principle

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay",amount)
}

type stripe struct{}

func (s stripe) pay(amount float32){
	fmt.Println("Making payment using stripe : ",amount)
}

// interfaces

type paymenter interface{
	pay(amount float32)
}

func main() {
	// stripePaymentGw := stripe{}
	razorpayPaymentGw := razorpay{}
	newPayment := payment{
		gateway: razorpayPaymentGw,
	}

	newPayment.makePayment(100)
}