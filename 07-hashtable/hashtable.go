package hashtable

import (
	"fmt"
	"strconv"
)

type HashTable[T any] struct {
	values []T
	length int
	size   int
}

func NewHashTable[T any](size int) *HashTable[T] {
	return &HashTable[T]{
		values: make([]T, size),
		length: 0,
		size:   size,
	}
}

func (ht *HashTable[T]) Add(key string, value T) error {
	hashKey := ht.hash(key)
	if hashKey >= ht.size {
		return fmt.Errorf("key %v is out of range", key)
	}
	ht.values[ht.hash(key)] = value
	ht.length++
	return nil
}

func (ht *HashTable[T]) hash(key string) int {
	if ht.size <= 0 {
		return 0
	}
	value, _ := strconv.ParseInt(key, 36, 0)
	//fmt.Printf("%v -> %v -> %v\n", key, value, value%int64(ht.size))
	return int(value % int64(ht.size))
}

func (ht *HashTable[T]) Get(s string) T {
	return ht.values[ht.hash(s)]
}
