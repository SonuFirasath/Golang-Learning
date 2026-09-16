package main

import (
	"fmt"
)

func main() {
	// simple switch

	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// case 4:
	// 	fmt.Println("Four")
	// case 5:
	// 	fmt.Println("Five")
	// default:
	// 	fmt.Println("No match found")
	// }

	// Multiple condition switch

	// switch time.Now().Weekday() {

	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's weekend")
	// default:
	// 	fmt.Println("It's not a weekend.")
	// }

	// type switch

	whoAmI := func(i interface{}){
		switch t:= i.(type) {
		case int:
			fmt.Println("It's an integer")
		case float64:
			fmt.Println("It's a float")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("It's a boolian")
		default:
			fmt.Println("This is not an identifed type",t)
		}
	}

	whoAmI(80)
}