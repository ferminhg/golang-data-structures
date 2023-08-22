package heaps

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type AllowedData interface {
	constraints.Integer
}

type MaxHeap[T AllowedData] struct {
	heap     []T
	elements int
}

func NewMaxHeap[T AllowedData]() *MaxHeap[T] {
	return &MaxHeap[T]{
		heap:     make([]T, 0),
		elements: 0,
	}
}

func (h *MaxHeap[T]) Insert(value T) {
	if len(h.heap) >= h.elements {
		h.elements++
		h.heap = append(h.heap, value)
		h.percolateUp(len(h.heap) - 1)
		return
	}
	h.heap[h.elements] = value
	h.elements++
	h.percolateUp(h.elements - 1)
}

func (h *MaxHeap[T]) percolateUp(index int) {
	if index <= 0 {
		return
	}
	parentIndex := (index - 1) / 2
	if h.heap[parentIndex] < h.heap[index] {
		h.heap[parentIndex], h.heap[index] = h.heap[index], h.heap[parentIndex]
		h.percolateUp(parentIndex)
	}
}

func (h *MaxHeap[T]) maxHeapify(index int) {
	if index >= h.elements {
		return
	}
	leftIndex := index*2 + 1
	rightIndex := index*2 + 2
	largest := index
	switch {
	case leftIndex < h.elements && h.heap[leftIndex] > h.heap[largest]:
		largest = leftIndex
	case rightIndex < h.elements && h.heap[rightIndex] > h.heap[largest]:
		largest = rightIndex
	}
	if largest != index {
		h.heap[largest], h.heap[index] = h.heap[index], h.heap[largest]
		h.maxHeapify(largest)
	}
}

func (h *MaxHeap[T]) GetMax() (T, bool) {
	if h.elements == 0 {
		return *new(T), false
	}
	return h.heap[0], true
}

func (h *MaxHeap[T]) RemoveMax() (T, bool) {
	if h.elements == 0 {
		return *new(T), false
	}
	result := h.heap[0]
	h.heap[0] = h.heap[h.elements-1]
	h.elements--
	h.maxHeapify(0)
	return result, true
}

func (h *MaxHeap[T]) Println() string {
	return fmt.Sprintf(
		"Heap:\nElements: %v\n %v",
		h.elements,
		h.heap,
	)
}
