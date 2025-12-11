package educativedsa

import "fmt"

type CorrectMinStack struct {
	MaxLength    int
	mainStack    SortedInsertStack
	minimumStack SortedInsertStack
}

// pop()

func (m *CorrectMinStack) Pop() int {
	m.minimumStack.Pop()

	value := m.mainStack.Peek()
	m.mainStack.Pop()
	return value

}

// push

func (m *CorrectMinStack) Push(value int) {
	m.mainStack.Push(value)
	fmt.Println("dfdnfd", m.minimumStack.Peek(), value)
	if !m.minimumStack.IsEmpty() && m.minimumStack.Peek() < value {
		m.minimumStack.Push(m.minimumStack.Peek())
	} else {
		m.minimumStack.Push(value)
	}
}

// Min

func (m *CorrectMinStack) Min() int {

	return m.minimumStack.Peek()
}
