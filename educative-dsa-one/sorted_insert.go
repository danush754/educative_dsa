package educativedsa

import "fmt"

type SortedInsertStack struct {
	data []int
}

// isEmpty()

func (s *SortedInsertStack) IsEmpty() bool {

	return len(s.data) == 0

}

// length()

func (s *SortedInsertStack) Length() int {

	return len(s.data)
}

// print()

func (s *SortedInsertStack) Print() {

	for i := 0; i < len(s.data); i++ {
		fmt.Print(s.data[i], " ")
	}

	fmt.Println()

}

// push()

func (s *SortedInsertStack) Push(data int) {

	s.data = append(s.data, data)
}

// pop()

func (s *SortedInsertStack) Pop() int {

	if !s.IsEmpty() {

		data := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return data

	}

	return -1
}

// peek()

func (s *SortedInsertStack) Peek() int {

	if !s.IsEmpty() {

		return s.data[len(s.data)-1]
	}

	return -1
}

func SortedInsert(stk *SortedInsertStack, element int) {

}
