package main

import "fmt"

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack struct{
	elements []int
}

func main() {
	// nums := []int{1,2,3,4}
	names := []string{"golang","typescript"}
	// priceStringSlice(names)
	printSlice(names)
}