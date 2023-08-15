package trees

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BinaryTree(t *testing.T) {

	t.Run("Given an element BinaryTree when insert on binary tree Then we can find the element ", func(t *testing.T) {
		tree := NewBinaryTree[int]()
		tree.Insert(10)

		assert.True(t, tree.Find(10))
		assert.False(t, tree.Find(20))
	})

	t.Run("Given an full BinaryTree when find element Then we can find the element ", func(t *testing.T) {

		tree := NewBinaryTree[int]()
		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(15)

		assert.True(t, tree.Find(10))
		assert.True(t, tree.Find(5))
		assert.True(t, tree.Find(15))
		assert.False(t, tree.Find(20))
	})

	t.Run("Given an full BinaryTree  when find element Then we can find the element ", func(t *testing.T) {

		tree := NewBinaryTree[int]()
		tree.Insert(10)
		tree.Insert(5)
		tree.Insert(15)
		tree.Insert(20)

		fmt.Println(tree.Print())

		assert.True(t, tree.Find(10))
		assert.True(t, tree.Find(5))
		assert.True(t, tree.Find(15))
		assert.True(t, tree.Find(20))
		assert.Equal(t, " 5 10 15 20", tree.Print())
	})
}
