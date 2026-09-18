package main

import (
	"fmt"
)

// Slice
// Most used construct in go
//
func main(){
	// Uninitialized slice is nill
	// var nums []int

	// fmt.Println(nums==nil)

	// var nums = make([]int,0,5)
	// capacity -> Maximum number of elements can fit
	// fmt.Println(cap(nums))
	// fmt.Println(nums)

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)

	// fmt.Println(cap(nums))

	// fmt.Println(nums)

	// nums := []int{}

	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)

	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// var nums = make([]int, 0,5)
	// nums = append(nums, 2)
	// var nums2 = make([]int , len(nums))

	// copy function

	// copy(nums2,nums)

	// fmt.Println(nums,nums2)


	// slice operator

	// var nums = []int{1,2,3,4,5}

	// fmt.Println(nums[0:2])
	// fmt.Println(nums[:2])
	// fmt.Println(nums[1:])

	// Slice package

	// var nums1 = []int{1,2}
	// var nums2 = []int{1,2}

	// fmt.Println(slices.Equal(nums1,nums2))

	var nums = [][]int{{1,2,3},{4,5,6}}

	fmt.Println(nums)

}