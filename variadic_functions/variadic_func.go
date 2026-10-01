package main

import "fmt"

// func sum(nums ...interface{} --> we can use interface{} which can accept any type of parameters but again we can't perform addition operation because we can't do addition in all types of parameters) int{
func sum(nums ...int) int{
	total := 0

	for _,num:= range nums{
		total = total+num
	}
	return total
}

func main() {
	// fmt.Println(1,2,3,4,5,5,66) // ---> this is a variadic function in go (println) because it can accept any n number of parameters.

	// fmt.Println(sum(4,3,1,2,3,4))

	// if i have values in slice then how to paas it to sum variadic func

	nums := []int{1,2,3,4,5}

	result := sum(nums...)

	fmt.Println(result)
}