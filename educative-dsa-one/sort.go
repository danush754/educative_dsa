package educativedsa

func BubbleSort(array []int, comp func(val1 int, val2 int) bool) {

	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {

			if comp(array[j], array[j+1]) {

				array[j], array[j+1] = array[j+1], array[j]

			}
		}
	}

}

func IsGreater(val1, val2 int) bool {
	return val1 > val2
}

func InsertionSort(nums []int) {

	for i := 1; i < len(nums); i++ {

		temp := nums[i]

		j := i - 1

		for j >= 0 && nums[j] > temp {
			nums[j+1] = nums[j]
			j = j - 1
		}

		nums[j+1] = temp
	}
}

func SelectionSort(nums []int) {

	for i := 0; i < len(nums); i++ {

		for j := i; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
