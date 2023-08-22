package hashtable

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_HashTable(t *testing.T) {
	t.Run("Given a HashTable, When I add a value, Then I can retrieve it", func(t *testing.T) {
		ht := NewHashTable[int](10)
		err := ht.Add("1", 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, ht.Get("1"))
	})

	t.Run("Given a HashTable, When I add two values with the same key, Then the second value overwrites the first", func(t *testing.T) {
		ht := NewHashTable[int](10)
		ht.Add("a", 1)
		ht.Add("a", 2)
		assert.Equal(t, 2, ht.Get("a"))

		ht.Add("b", 3)
		assert.Equal(t, 2, ht.Get("a"))
		assert.Equal(t, 3, ht.Get("b"))
	})

	t.Run("Given a HashTable, When I add a value with a key that is out of range, Then I get an error", func(t *testing.T) {
		ht := NewHashTable[int](0)
		err := ht.Add("1", 1)
		require.Error(t, err)
		assert.Equal(t, "key 1 is out of range", err.Error())
	})
}
