package heaps

import (
	objectmother "github.com/ferminhg/golang-data-structures/internal/object-mother"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MaxHeap(t *testing.T) {
	intMO := objectmother.NewIntMotherObject()
	t.Run("Given an empty MaxHeap, Insert should add the value to the heap", func(t *testing.T) {
		heap := NewMaxHeap[int]()
		expected := intMO.Get()
		heap.Insert(expected)
		result, ok := heap.GetMax()
		assert.True(t, ok)
		assert.Equal(t, expected, result)
	})

	t.Run("Given a MaxHeap with two elements, GetMax should return the largest element", func(t *testing.T) {
		heap := NewMaxHeap[int]()
		expected := intMO.Get()
		heap.Insert(expected)
		heap.Insert(expected - 1)
		//fmt.Println(heap.Println())
		result, ok := heap.GetMax()
		assert.True(t, ok)
		assert.Equal(t, expected, result)
	})

	t.Run("Given a MaxHeap with five elements, GetMax should return the largest element", func(t *testing.T) {
		heap := NewMaxHeap[int]()
		expected := intMO.Get()
		heap.Insert(expected)
		heap.Insert(expected - 1)
		heap.Insert(expected + 1)
		heap.Insert(expected + 2)
		heap.Insert(expected + 3)
		//fmt.Println(heap.Println())
		result, ok := heap.GetMax()
		assert.True(t, ok)
		assert.Equal(t, expected+3, result)
	})

	t.Run("Given a MaxHeap with one element When RemoveMax is called, Then the element is removed", func(t *testing.T) {
		heap := NewMaxHeap[int]()
		expected := intMO.Get()
		heap.Insert(expected)
		result, ok := heap.RemoveMax()
		assert.True(t, ok)
		assert.Equal(t, expected, result)
	})

	t.Run("Given a MaxHeap with five elements When RemoveMax is called Then should return elements in order", func(t *testing.T) {
		heap := NewMaxHeap[int]()
		expected := intMO.Get()
		heap.Insert(expected)
		heap.Insert(expected - 1)
		heap.Insert(expected + 1)
		heap.Insert(expected + 2)
		heap.Insert(expected + 3)
		result, _ := heap.RemoveMax()
		assert.Equal(t, expected+3, result)
		result, _ = heap.RemoveMax()
		assert.Equal(t, expected+2, result)
		result, _ = heap.RemoveMax()

		assert.Equal(t, expected+1, result)
		result, _ = heap.RemoveMax()
		assert.Equal(t, expected, result)
		result, _ = heap.RemoveMax()
		assert.Equal(t, expected-1, result)
		_, ok := heap.RemoveMax()
		assert.False(t, ok)
	})

}
