package stacks

import (
	object_mother "github.com/ferminhg/golang-data-structures/internal/object-mother"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Stack(t *testing.T) {
	intFaker := object_mother.NewIntMotherObject()
	t.Run("Given an element When Push on stack Then the Size is one", func(t *testing.T) {
		s := NewLIFO[int](MaxSize)

		err := s.Push(intFaker.Get())
		require.NoError(t, err)
		assert.Equal(t, 1, s.Size())
	})

	t.Run("Given an empty stack When Pop on stack Then the error is not nil", func(t *testing.T) {
		s := NewLIFO[int](MaxSize)
		_, err := s.Pop()
		require.Error(t, err)
	})

	t.Run("Given a stack with one element When Pop on stack Then the error is nil", func(t *testing.T) {
		s := NewLIFO[int](MaxSize)
		randomValue := intFaker.Get()
		err := s.Push(randomValue)
		require.NoError(t, err)
		assert.Equal(t, 1, s.Size())
		expected, _ := s.Pop()
		assert.Equal(t, randomValue, expected)
		assert.Equal(t, 0, s.Size())
	})

	t.Run("Given a stack with one max size When Push on stack Then the error is not nil", func(t *testing.T) {
		s := NewLIFO[int](1)
		err := s.Push(intFaker.Get())
		require.NoError(t, err)
		err = s.Push(intFaker.Get())
		require.Error(t, err)
	})
}
