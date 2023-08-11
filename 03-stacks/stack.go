package stacks

import "errors"

type Stacks[T any] interface {
	Pop() (T, error)
	Push(value T)
	Size() int
}

type StackLIFO[T any] struct {
	maxSize int
	top     int
	items   []T
}

const MaxSize = 100

func NewLIFO[T any](size int) *StackLIFO[T] {
	return &StackLIFO[T]{
		maxSize: size,
		top:     -1,
		items:   make([]T, size),
	}
}

func (s *StackLIFO[T]) Pop() (T, error) {
	if s.top == -1 {
		return *new(T), errors.New("stack is empty")
	}
	item := s.items[s.top]

	s.items[s.top] = *new(T)
	s.top--
	return item, nil
}

func (s *StackLIFO[T]) Push(value T) error {
	if s.top == s.maxSize-1 {
		return errors.New("stack is full")
	}

	s.top++
	s.items[s.top] = value
	return nil
}

func (s *StackLIFO[T]) Size() int {
	return s.top + 1
}
