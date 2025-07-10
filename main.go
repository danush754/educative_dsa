package main

import (
	educativedsa "github.com/danush754/educative_dsa/educative-dsa-one"
)

func main() {
	// var data []int = []int{1, 4, 6, 8, 10}
	// var key = 4

	var array []int = []int{1}
	var k = 2

	// fmt.Println(educativedsa.SequentialSearch(data, key))
	// fmt.Println(educativedsa.BinarySearch(data, key))
	// fmt.Println(educativedsa.LargestSubarraySum_1(array))
	// fmt.Println(educativedsa.LargestSubarraySum_Kandane(array))
	// fmt.Println(educativedsa.RotateElementsByK(array, k))
	// educativedsa.RotateElementsByK_two(array, k)
	educativedsa.RotateElementsByK_three(array, k)

}
