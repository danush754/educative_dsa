package educativedsa

func MergeSort(nums []int) {
	size := len(nums)
	tempArr := make([]int, size)
	mergeStr(nums, tempArr, 0, size-1)

}

func mergeStr(arr, tempArr []int, lowerindex, upperIndex int) {

	// if upperIndex == lowerindex we will return i.e
	//  the array is of length 1 eg: [1]

	if lowerindex >= upperIndex {
		return
	}

	middleIndex := (lowerindex + upperIndex) / 2

	mergeStr(arr, tempArr, lowerindex, middleIndex)
	mergeStr(arr, tempArr, middleIndex+1, upperIndex)
	merge(arr, tempArr, lowerindex, upperIndex, middleIndex)

}

func merge(arr, tempArr []int, lowerIndex, upperIndex, middleIndex int) {
	lowerStart := lowerIndex
	lowerStop := middleIndex
	upperStart := middleIndex + 1
	upperStop := upperIndex

	count := lowerIndex

	for lowerStart <= lowerStop && upperStart <= upperStop {
		if arr[lowerStart] <= arr[upperStart] {
			tempArr[count] = arr[lowerStart]
			lowerStart++
		} else {
			tempArr[count] = arr[upperStart]
			upperStart++
		}
		count++
	}

	for lowerStart <= lowerStop {
		tempArr[count] = arr[lowerStart]
		lowerStart++
		count++
	}

	for upperStart <= upperStop {
		tempArr[count] = arr[upperStart]
		upperStart++
		count++
	}

	for i := lowerIndex; i <= upperIndex; i++ {
		arr[i] = tempArr[i]
	}

}
