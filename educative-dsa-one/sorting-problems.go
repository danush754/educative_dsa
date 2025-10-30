package educativedsa

import (
	"fmt"
	"math"
)

func FindmaxAndMin(arr []int) (min, max int) {

	min = math.MaxInt64
	max = math.MinInt64

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}

		if arr[i] > max {
			max = arr[i]
		}
	}

	return min, max
}

func SortArr(arr []int, maxVal, minVal int) {
	rnge := maxVal - minVal
	count := make([]int, rnge)

	for i := 0; i < len(arr); i++ {
		count[arr[i]-minVal]++
	}

	j := 0
	for i := 0; i < rnge; i++ {
		for ; count[i] > 0; count[i]-- {
			arr[j] = i + minVal
			j++
		}
	}
}

func PartionOnesAndZero(arr []int) {
	leftPointer := 0
	rightPointer := len(arr) - 1

	fmt.Println("skssk", leftPointer, rightPointer)

	for leftPointer < rightPointer {
		fmt.Println("sjssj", arr[leftPointer], arr[rightPointer])

		for arr[leftPointer] == 0 {
			leftPointer++
		}

		for arr[rightPointer] == 1 {
			rightPointer--
		}

		if leftPointer < rightPointer {
			arr[leftPointer], arr[rightPointer] = Swap(arr[leftPointer], arr[rightPointer])
		}

	}
}

func Swap(x, y int) (int, int) {
	return y, x
}

func Partition012(arr []int) {

	left := 0
	right := len(arr) - 1
	index := 0

	fmt.Println("left", left, right)
	for index <= right {
		if arr[index] == 0 {
			arr[index], arr[left] = arr[left], arr[index]
			left++
			index++

		} else if arr[index] == 2 {
			arr[index], arr[right] = arr[right], arr[index]
			right--
		} else {
			index++
		}

	}
}
