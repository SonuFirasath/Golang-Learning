package main

import (
	"fmt"
)

//Sending data to goroutine
// func processNum(num chan int){
// 	fmt.Println("processing number",<-num)
// }

// receiving data from goroutine

func sum(result chan int, num1 int,num2 int){
	sumResult := num1 + num2

	result <- sumResult
}
// goroutine synchronizer
func task(done chan bool){
	defer func(){done<-true}()
	fmt.Println("processing ....")
}

func emailSender(emailChan <-chan string,done chan<- bool){
	defer func ()  {
		done <- true
	}()
	for email := range emailChan{
		fmt.Println("sendin email to ", email)
	}
}


func main() {

	// emailChan := make(chan string,100) // ---> buffer channel, the 100 is the size of data that we can send in this channel so that it doesn't act like blocking channel

	// done := make(chan bool)

	// go emailSender(emailChan,done)

	// for i:=0;i<5;i++{
	// 	emailChan <- fmt.Sprintf("%d@gamil.com",i)
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println("done printing")

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func(){
		chan1 <- 10
	}()

	go func(){
		chan2 <- "golang"
	}()

	for i:=0; i<2; i++{
		select{
		case chan1Val := <- chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <- chan2:
			fmt.Println("received data from chan2",chan2Val)
		}
	}

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// close(emailChan)
	// <- done

	// done := make(chan bool)

	// go task(done)

	// <- done //block

	// result := make(chan int)

	// go sum(result,3,3)

	// res := <- result

	// fmt.Print("The sum of these two numbers is : ",res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// numChan <- 5

	// time.Sleep(time.Second*2)

	// messageChan := make(chan string)

	// messageChan <- "ping" // channels are blocking

	// msg := <-messageChan

	// fmt.Println(msg)
}