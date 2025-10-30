package educativedsa

import "fmt"

func QuickSelect(arr []int, lower int, upper int, k int) int {
	// we are checking here whether k is greater than 0 and lesser than the size of
	
	if k > 0 && k <= upper-lower+1 {
		pivot := arr[upper] // last element as pivot
		start := lower
		stop := lower

		fmt.Println("pivvvovv", pivot, start, stop)
		for stop <= upper-1 { // partitioning
			if arr[stop] <= pivot {
				swap(arr, start, stop)
				start++
			}
			stop++
		}

		swap(arr, start, upper)
		pivotindex := start
		if pivotindex-lower == k-1 {
			return arr[pivotindex] // return kth element
		}
		if pivotindex-lower > k-1 {
			return QuickSelect(arr, lower, pivotindex-1, k) // find element in left subarray
		} else {
			return QuickSelect(arr, pivotindex+1, upper, k-pivotindex+lower-1) // find element in right subarray
		}
	}
	fmt.Println("The value of k is out of range")
	return -1 // returning arbitrary value
}
func swap(arr []int, first int, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}
