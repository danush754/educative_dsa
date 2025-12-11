package educativedsa

import (
	"fmt"
)

type SortedInsertStack struct {
	Data []int
}

// isEmpty()

func (s *SortedInsertStack) IsEmpty() bool {

	return len(s.Data) == 0

}

// length()

func (s *SortedInsertStack) Length() int {

	return len(s.Data)
}

// print()

func (s *SortedInsertStack) Print() {

	for i := 0; i < len(s.Data); i++ {
		fmt.Print(s.Data[i], " ")
	}

	fmt.Println()

}

// push()

func (s *SortedInsertStack) Push(Data int) {

	s.Data = append(s.Data, Data)
}

// pop()

func (s *SortedInsertStack) Pop() int {

	if !s.IsEmpty() {

		Data := s.Data[len(s.Data)-1]
		s.Data = s.Data[:len(s.Data)-1]
		return Data

	}

	return -1
}

// peek()

func (s *SortedInsertStack) Peek() int {

	if !s.IsEmpty() {

		return s.Data[len(s.Data)-1]
	}

	return -1
}

func (stk *SortedInsertStack) SortedInsert(element int) {
	var temp int

	if stk.Length() == 0 || stk.Peek() < element {
		stk.Push(element)
	} else {
		temp = stk.Pop()
		stk.SortedInsert(element)
		stk.Push(temp)
	}
}

func (stk *SortedInsertStack) SortStack() {

	var temp int

	if stk.Length() == 0 {
		return
	}

	temp = stk.Pop()
	stk.SortStack()
	stk.SortedInsert(temp)

}

func (stk *SortedInsertStack) BottomInsert(element int) {
	var temp int
	if stk.Length() == 0 {
		stk.Push(element)
	} else {
		temp = stk.Pop()
		stk.BottomInsert(element)
		stk.Push(temp)
	}
}

func (stk *SortedInsertStack) ReverseStack() {
	var temp int

	if stk.Length() > 0 {

		temp = stk.Pop()
		stk.ReverseStack()

		fmt.Println("temp", temp)
		stk.BottomInsert(temp)
	}
}

func (stk *SortedInsertStack) ReverseKelements(count int) {

	newQueue := new(Queue)

	i := 0

	for stk.Length() > 0 && i < count {
		newQueue.Enqueue(stk.Pop())
		i++
	}

	for newQueue.Length() != 0 {
		stk.Push(newQueue.Dequeue().(int))
	}

	fmt.Printf("newQueue: %v\n", newQueue)

}
