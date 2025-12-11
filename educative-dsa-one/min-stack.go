package educativedsa

import (
	"fmt"
	"math"
)

type MinStack struct {
	minValue int
	data     []int
}

func (m *MinStack) IsEmpty() bool {
	return len(m.data) == 0
}

func (m *MinStack) Length() int {
	return len(m.data)
}

func (m *MinStack) Push(value int) {
	if m.Length() == 0 {
		m.minValue = value
	} else {
		if value < m.minValue {
			m.minValue = value
		}
	}

	m.data = append(m.data, value)
}

func (m *MinStack) Pop() int {
	length := m.Length()

	value := m.data[length-1]
	m.data = m.data[:length-1]

	fmt.Println("from pop", m.minValue, value)

	if m.IsEmpty() {
		m.minValue = math.MinInt64
	} else if value == m.minValue {
		temp := math.MaxInt64
		for i := range m.data {
			fmt.Println("ss", m.data[i], temp)
			if m.data[i] < temp {
				m.minValue = m.data[i]
			}
		}
	}

	return value
}

func (m *MinStack) Min() int {

	return m.minValue

}
