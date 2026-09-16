package main

import "fmt"

func main() {
	// age := 18

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	// var role = "admin"
	// var hasPermissions = false

	// if role == "admin" && hasPermissions {
	// 	fmt.Println("Yes")
	// } else if role == "admin"{
	// 	fmt.Println("Yes from else if")
	// } else if hasPermissions{
	// 	fmt.Println("Yes because you have permission")
	// } else{
	// 	fmt.Println("No you don't have access")
	// }

	if age := 15; age >=18{
		fmt.Println("Person is an adult",age)
	} else if age >= 12 {
		fmt.Println("Person is teenager",age)
	}

	// Go does not have ternary operator as per 1.22 version, you will have to use normal if else

}