package main

import (
	"fmt"

	educativedsa "github.com/danush754/educative_dsa/educative-dsa-one"
)

type Book struct {
	Title  string `json:"title" genre:"fiction"`
	Author string `json:"author" genre:"non-fiction"`
}

// Define a root handler for requests to function homePage, and start the webserver combined with error-handling
func main() {

	// http.HandleFunc("/", educativedsa.HomePage)

	// http.ListenAndServe(":8080", nil)

	qc := []int{9, 1, 9, 10, 10, 7, 2, 8, 2, 7, 3, 6, 4, 5}
	size := len(qc)

	fmt.Println("qujjjs", educativedsa.QuickSelect(qc, 0, size-1, 2))

	var countSortData = []int{9, 1, 9, 10, 10, 7, 2, 8, 2, 7, 3, 6, 4, 5}

	minVal, maxVal := educativedsa.FindMinAndMax(countSortData)

	fmt.Println("minmax", minVal, maxVal)
	educativedsa.CountSort(countSortData, minVal, maxVal+1)

	// var zeroAndOnes = []int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1}

	var zerosAndTwos = []int{0, 0, 2, 0, 2, 1, 0, 1}

	// minval, maxval := educativedsa.FindmaxAndMin(zerosAndTwos)

	// educativedsa.SortArr(zerosAndTwos, maxval+1, minval)

	// educativedsa.PartionOnesAndZero(zeroAndOnes)

	educativedsa.Partition012(zerosAndTwos)

	fmt.Println("zeroAndOnes", zerosAndTwos)

	var sortedStack educativedsa.SortedInsertStack

	sortedStack.Push(1)
	sortedStack.Push(7)
	sortedStack.Push(8)
	sortedStack.Push(5)

	// sortedStack.SortedInsert(5)
	// sortedStack.SortStack()
	// sortedStack.BottomInsert(0)

	// sortedStack.ReverseStack()
	sortedStack.InsertElementAtk(3, 5)
	// sortedStack.ReverseKelements(2)
	sortedStack.Print()
}
