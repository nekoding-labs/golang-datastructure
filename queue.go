package main

import "errors"

type Queue struct {
	High []int
	Low  []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(val int, highPriority bool) Queue {

	if highPriority {
		q.High = append(q.High, val)
	} else {
		q.Low = append(q.Low, val)
	}

	return *q
}

func (q Queue) Size() int {
	return len(q.High) + len(q.Low)
}

func (q Queue) IsEmpty() bool {
	return q.Size() < 1
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	var firstElement int
	if len(q.High) > 1 {
		firstElement, q.High = q.High[0], q.High[1:]
		return firstElement, nil
	}

	firstElement, q.Low = q.Low[0], q.Low[1:]
	return firstElement, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	if len(q.High) > 0 {
		return q.High[0], nil
	}

	return q.Low[0], nil
}
