package educativedsa

import "fmt"

func SequentialSearch(data []int, key int) bool {

	for _, value := range data {
		if value == key {
			return true
		} else {
			continue
		}
	}

	return false
}

func BinarySearch(data []int, key int) bool {

	var leftPointer = 0
	var rightPointer = len(data) - 1
	var mid int

	for leftPointer <= rightPointer {
		mid = (leftPointer + rightPointer) / 2

		if data[mid] == key {
			return true
		} else {
			if data[mid] < key {
				leftPointer = mid + 1
			} else {
				rightPointer = mid - 1
			}
		}
	}

	return false
}

func LargestSubarraySum_1(array []int) int {

	var largestSum = 0

	for i := 0; i < len(array); i++ {
		var tempSum = 0

		for j := i; j < len(array); j++ {

			tempSum += array[j]

			if tempSum > largestSum {
				largestSum = tempSum
			}
		}
	}

	return largestSum
}

func LargestSubarraySum_Kandane(array []int) int {

	if len(array) == 0 {
		return 0
	}

	if len(array) == 1 {
		return array[0]
	}
	var maxVal = array[0]
	var tempMaxVal = array[0]

	for i := 1; i < len(array); i++ {

		tempMaxVal = max(array[i], tempMaxVal+array[i])

		maxVal = max(tempMaxVal, maxVal)

	}

	return maxVal
}

func max(a, b int) int {

	if a > b {
		return a
	}

	return b
}

func RotateElementsByK(array []int, k int) []int {

	for i := 0; i < k; i++ {
		first := array[len(array)-1]

		for j := len(array) - 1; j > 0; j-- {

			array[j] = array[j-1]

		}

		array[0] = first

	}

	return array

}

func RotateElementsByK_two(array []int, k int) {

	if len(array) == 0 || len(array) == 1 {
		return
	}

	k = k % len(array)
	fmt.Println("klklkl", k)

	var dupeArray = make([]int, len(array))

	var startval = 0
	for i := len(array) - k; i < len(array); i++ {

		dupeArray[startval] = array[i]
		startval++
	}

	startval = k

	for i := 0; i < len(array)-k; i++ {
		dupeArray[startval] = array[i]
		startval++
	}

	for i := 0; i < len(array); i++ {
		array[i] = dupeArray[i]
	}

}

func RotateElementsByK_three(array []int, k int) {

	if len(array) == 0 || len(array) == 1 {
		return
	}

	k = k % len(array)

	array = ReverseArray(array, 0, len(array)-1)

	array = ReverseArray(array, 0, k-1)

	array = ReverseArray(array, k, len(array)-1)

}

func ReverseArray(array []int, startVal, endVal int) []int {

	for startVal < endVal {
		array[startVal], array[endVal] = array[endVal], array[startVal]
		startVal++
		endVal--
	}

	return array
}
