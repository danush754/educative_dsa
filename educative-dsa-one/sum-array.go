package educativedsa

import (
	"fmt"
)

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

func Waveform(array []int) []int {

	if len(array) == 0 || len(array) == 1 {
		return array
	}

	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j+1] <= array[j] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	for i := 0; i < len(array)-1; i += 2 {
		if array[i+1] >= array[i] {
			array[i+1], array[i] = array[i], array[i+1]
		}
	}

	return array
}

func IndexArray(array []int) []int {

	temp := 0

	for i := 0; i < len(array); i++ {
		for array[i] != -1 && array[i] != i {
			temp = array[i]
			array[i] = array[temp]
			array[temp] = temp
		}

	}

	return array
}

func SortanArray(array []int) []int {
	size := len(array)
	var dupeArray = make([]int, size)

	for i := 0; i < len(array); i++ {
		temp := array[i]

		dupeArray[temp-1] = array[i]

	}

	for i := 0; i < len(array); i++ {
		array[i] = dupeArray[i]
	}

	return array

}

func BinarySort(array []int) []int {

	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}

		}

	}

	return array
}

func SmallestPositiveMissing(array []int) int {
	startVal := 1

	for _, val := range array {
		if val == startVal {
			startVal++
			continue

		} else {
			return startVal
		}
	}

	return -1
}

func MaximumMinimumArray(arr []int, size int) []int {

	var dupeArr = make([]int, size)

	leftPointer := 0
	rightPointer := size - 1

	for i := 0; i < size-1; i += 2 {

		dupeArr[i], dupeArr[i+1] = arr[rightPointer], arr[leftPointer]
		leftPointer++
		rightPointer--
	}

	if size%2 != 0 {
		dupeArr[size-1] = arr[leftPointer]
	}

	for i := 0; i < size; i++ {
		arr[i] = dupeArr[i]
	}

	return arr
}

func MaximumDiffOne(arr []int, size int) int {

	maxdifference := -1
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			fmt.Println("j,i", j, i)
			if j > i && arr[j] > arr[i] {
				fmt.Println("hhhs", j-i)
				if j-i > maxdifference {
					maxdifference = j - i
				}
			}

		}
	}

	return maxdifference
}

// func QuizArray(arr1, arr2 []int, size1, size2 int) int {

// 	i, j := 0, 0
// 	sum1, sum2, result :=

// 	return maxSum
// }
