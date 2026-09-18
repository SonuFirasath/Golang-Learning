package main

import "fmt"

// maps -> hash,object,dict
func main(){
	// creating map

	// m:= make(map[string]string) // map[string --> key type]string --> value type

	// Setting an element

	// m["name"] = "golang"
	// m["area"] = "backend"

	// Get an element

	// fmt.Println(m["name"],m["area"])
	// fmt.Println(m["phone"]) //---> does not exist in the map so empty string will be returned.

	// m := make(map[string]int)

	// m["age"] = 22
	// m["price"] = 50
	
	// fmt.Println(m)
	// delete(m,"price")
	
	// fmt.Println(m["phone"])
	
	// fmt.Println(len(m))
	// fmt.Println(m)

	m := map[string]int{"price": 40, "phone":3}

	fmt.Println(m)

	_,ok := m["price"]

	if ok{
		fmt.Println("all ok")
	} else{
		fmt.Println("not ok")
	}
}