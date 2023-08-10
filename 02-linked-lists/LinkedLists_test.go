package linked_lists

import (
	objectmother "github.com/ferminhg/golang-data-structures/internal/object-mother"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SinglyLinkedList(t *testing.T) {
	intFaker := objectmother.NewIntMotherObject()
	t.Run("should create a new linked list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		assert.Nil(t, list.head)
		assert.Equal(t, 0, list.size)
	})

	t.Run("should add a new node to the list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		randomValue := intFaker.Get()
		list.Add(randomValue)

		assert.Equal(t, 1, list.size)
		assert.Equal(t, randomValue, list.head.value)
	})

	t.Run("should return the length of the list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		randomSize := intFaker.Get()

		for i := 0; i < randomSize; i++ {
			list.Add(intFaker.Get())
		}

		assert.Equal(t, randomSize, list.Length())
	})

	t.Run("should return the last node of the list", func(t *testing.T) {
		list := SinglyLinkedList[int]{}

		lastValue := intFaker.Get()
		list.Add(intFaker.Get())
		list.Add(intFaker.Get())
		list.Add(lastValue)

		lastNode := list.Last()
		assert.Equal(t, lastValue, lastNode.value)
	})

	t.Run("should return the first node of the list", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()

		firstValue := intFaker.Get()
		list.Add(firstValue)
		list.Add(intFaker.Get())
		list.Add(intFaker.Get())

		firstNode := list.First()
		assert.Equal(t, firstValue, firstNode.value)
	})

	t.Run("should return the node at the given index", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		secondNodeValue := intFaker.Get()
		list.Add(intFaker.Get())
		list.Add(secondNodeValue)
		list.Add(intFaker.Get())

		node := list.Get(1)
		assert.Equal(t, secondNodeValue, node.value)
	})
}
