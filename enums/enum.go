package main

import "fmt"

// enumerated types

type orderStatus string

const (
	received  orderStatus = "received"
	confirmed orderStatus = "confirmed"
	prepared  orderStatus = "prepared"
	delivered orderStatus = "delivered"
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("changing order status to : ", status)
}

func main() {

	changeOrderStatus(received)

}
