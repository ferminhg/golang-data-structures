package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// Learning Go book exercise

func TestTooRigid(t *testing.T) {
	var x = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x)
}

func TestBasicSlides(t *testing.T) {
	var x = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x)

	assert.Zero(t, len([]int{}))
	assert.Equal(t, 12, len(x))

	var y []int
	y = append(y, 12)
	reflect.DeepEqual(y, [1]int{12})

	y = append(y, 12, 12, 12)
	assert.Equal(t, 4, len(y))

	y = append(y, x...)
	assert.Equal(t, 16, len(y))

	var c []int
	fmt.Println(len(c), cap(c))
	c = append(c, 10)
	fmt.Println(len(c), cap(c))
	c = append(c, 20)
	fmt.Println(len(c), cap(c))
	c = append(c, 30)
	fmt.Println(len(c), cap(c))
	c = append(c, 40)
	fmt.Println(len(c), cap(c))
	assert.Equal(t, 4, len(c))
	assert.Equal(t, 4, cap(c))
	c = append(c, 50)
	fmt.Println(len(c), cap(c))
	assert.Equal(t, 8, cap(c))

	x2 := []int{1, 2, 3, 4}
	y2 := x2[:2]
	z2 := x2[1:]

	x2[1] = 20
	y2[0] = 10
	z2[1] = 30

	reflect.DeepEqual(x2, []int{1, 20, 30, 4})

	a := []int{1, 2, 3, 4}
	fmt.Println(a[:3])
	fmt.Println(a[1:])
	num := copy(a[:3], a[1:])

	fmt.Println(a, num)
}

func TestMapsSets(t *testing.T) {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10, 100}

	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

}

func TestStructSets(t *testing.T) {
	intSet := map[int]struct{}{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10, 100}

	for _, v := range vals {
		intSet[v] = struct{}{}
	}

	fmt.Println(len(vals), len(intSet))
	if _, ok := intSet[100]; ok {
		fmt.Println("100 is in the set")
	}
}
