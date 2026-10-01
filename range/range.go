package main

import (
	"fmt"
)

// range is used in iteration over the data structures more oftenly

func main(){
	// nums := []int {6,7,8}

	// for i:=0; i<len(nums);i++{
	// 	fmt.Println(nums[i])
	// }
	// sum := 0
	// for i,num := range nums{
	// 	// sum = sum + num
	// 	fmt.Println(num,i)
	// }
	// fmt.Println(sum)


	// m := map[string]string{"fname":"jhon","lname":"doe"}

	// for k,v := range m{
	// 	fmt.Println(k,v)
	// }

	// uni code point rune -> c value 
	// i here is not exactly a index here it is a starting byte of rune
	// if it's less than 255 then it will fit in 1 byte, if more than that it might take more than 2 byte for that.
	// to print the value of each uni code point rune you have to use string(c) to print the value of that .

	for i,c := range "golang"{
		fmt.Println(i,string(c))
	}

}