package educativedsa

import (
	"fmt"
	"math"
)

func FindMinAndMax(data []int) (min, max int) {

	min = math.MaxInt64
	max = math.MinInt64

	for i := 0; i < len(data); i++ {
		if data[i] < min {
			min = data[i]
		}
		if data[i] > max {
			max = data[i]
		}
	}

	return min, max

}

func CountSort(arr []int, minVal, maxVal int) {

	rng := maxVal - minVal

	fmt.Println("eng", rng)

	count := make([]int, rng)
	fmt.Println("count", count)

	for i := 0; i < len(arr); i++ {
		count[arr[i]-minVal]++
	}

	fmt.Println("count", count)

	j := 0

	for i := 0; i < rng; i++ {
		for ; count[i] > 0; count[i]-- {
			arr[j] = i + minVal
			j++
		}
	}

	fmt.Println("arr", arr)
}
