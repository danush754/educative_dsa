package educativedsa

import (
	"fmt"
)

const capacity = 100

type Queue struct {
	Size  int
	Data  [capacity]interface{}
	Front int
	Back  int
}

// enqueue
func (q *Queue) Enqueue(value interface{}) {
	if q.Size >= capacity {
		return
	}

	q.Size++
	q.Data[q.Back] = value
	q.Back = (q.Back + 1) % capacity
	fmt.Println("ack", q.Back, q.Front)
}

// dequeue
func (q *Queue) Dequeue() interface{} {

	var value interface{}

	if q.Size <= 0 {
		return 0
	}

	q.Size--
	value = q.Data[q.Front]
	q.Front = (q.Front + 1) % 100
	fmt.Println("ack", q.Back, q.Front)
	return value
}

// isEmpty

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}

// length

func (q *Queue) Length() int {

	return q.Size
}
