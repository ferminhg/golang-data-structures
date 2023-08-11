package stacks

import (
	object_mother "github.com/ferminhg/golang-data-structures/internal/object-mother"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Queue(t *testing.T) {
	intFaker := object_mother.NewIntMotherObject()
	t.Run("Given a Queue When add value Then top return the value", func(t *testing.T) {
		queue := NewQueue[int]()
		randomValue := intFaker.Get()

		queue.Enqueue(randomValue)
		expected, err := queue.DeQueue()
		require.NoError(t, err)
		assert.Equal(t, randomValue, expected)

		_, err = queue.DeQueue()
		require.Error(t, err)
	})

	t.Run("Given a Queue When add values Then top return the values in order", func(t *testing.T) {
		queue := NewQueue[int]()
		randomValue1 := intFaker.Get()
		randomValue2 := intFaker.Get()

		queue.Enqueue(randomValue1)
		queue.Enqueue(randomValue2)
		expected1, err := queue.DeQueue()
		require.NoError(t, err)
		assert.Equal(t, randomValue1, expected1)

		expected2, err := queue.DeQueue()
		require.NoError(t, err)
		assert.Equal(t, randomValue2, expected2)

		_, err = queue.DeQueue()
		require.Error(t, err)
	})
}
