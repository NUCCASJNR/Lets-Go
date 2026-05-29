package main

import "fmt"

func slices() {
	nums := []int{10, 20, 30}
	nums = append(nums, 40)

	fmt.Println(nums)
}