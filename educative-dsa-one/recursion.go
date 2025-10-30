package educativedsa

import (
	"container/list"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

type Dhanush struct {
	Name string
	Age  int
}

func Factorial(i int) int {

	if i <= 1 {
		return 1
	}

	return i * Factorial(i-1)
}

func IntToHex(num int) string {

	if num == 0 {

		return ""
	}

	hexValues := "0123456789ABCDEF"

	return IntToHex(num/16) + string(hexValues[num%16])
}

func GCD(num1, num2 int) int {

	var largeNum, smallNum int

	if num1 > num2 {
		largeNum = num1
		smallNum = num2
	} else {
		largeNum = num2
		smallNum = num1
	}

	if largeNum%smallNum == 0 {
		return smallNum
	}

	return GCD(smallNum, largeNum%smallNum)
}

func Fibonacci(n int) int {

	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Permutation(data []int, i int, length int) {

	// base condition DO NOT ALTER IT
	fmt.Println("length", length)
	fmt.Println("i", i)
	if length-1 < i {
		temp := "{"
		fmt.Println("tempp", temp)
		for k := 0; k < length; k++ {
			temp += strconv.Itoa(data[k])
			temp += " "
		}
		temp += "} "
		fmt.Print(temp)
		return
	}
	// Write your code here

	Permutation(data, i+1, data[length])
}

func CheckDeleteInMaps() {
	justFloat := make(map[int]float32)

	justFloat[1] = 1.0
	justFloat[2] = 2.0

	fmt.Println("justFloat", justFloat)

	for key := range justFloat {
		if key == 2 {
			delete(justFloat, key)
		}
	}

	fmt.Println("justFloat", justFloat)
}

func CreateLinkedList(n int) *list.List {
	newList := list.New()

	for i := 1; i <= n; i++ {
		newList.PushBack(i)
	}

	fmt.Println("memory", unsafe.Sizeof(Dhanush{}))

	return newList
}

func AnalyseStructTags(s reflect.Type) {

	fieldCount := s.NumField()

	for i := 0; i < fieldCount; i++ {
		field := s.Field(i)
		fmt.Println(field.Tag)
		fmt.Println(field.Tag.Get("json"))
		fmt.Println(field.Tag.Get("genre"))
	}
}
