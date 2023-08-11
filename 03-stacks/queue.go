package stacks

import "errors"

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

func (q *Queue[T]) DeQueue() (T, error) {
	if len(q.items) == 0 {
		return *new(T), errors.New("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]

	return item, nil
}
