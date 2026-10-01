package main

import (
	"fmt"
	"time"
)

// order struct

type customer struct{
	name string
	phone string
}

type order struct{
	id string
	amount float32
	status string
	createdAt time.Time // nanosecond precision ---> fully accurate
	customer
}

func newOrder(id string, amount float32, status string) *order{
	// initial setup goes here....
	myOrder := order{
		id: id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount 
}

func main(){

	// myOrder := newOrder("1",40.30,"received")

	// fmt.Println(myOrder.amount)

	customers := customer{
		name: "Firasath",
		phone: "032993838392839",
	}

	orders := order{
		id: "3",
		amount: 45.00,
		status: "received",
		customer: customers,
	}

	fmt.Println(orders)

	language := struct {
		name string
		isGood bool
	}{"golang",true}

	fmt.Println(language.name)

	// If you don't set any field to struct object, the default valu is zero value for that feild.
	// int -> 0, float -> 0, string -> "", bool -> false
	// myOrder := order{
	// 	id: "1",
	// 	amount: 50.0,
	// 	status: "received",
	// }

	// myOrder.createdAt = time.Now()
	// myOrder.changeStatus("confirmed")
	// myOrder2 := order{
	// 	id: "2",
	// 	amount: 100,
	// 	status: "pending",
	// 	createdAt: time.Now(),
	// }

	// fmt.Println(myOrder.status,myOrder.getAmount())

	// fmt.Println(myOrder)
	// fmt.Println(myOrder2)
}