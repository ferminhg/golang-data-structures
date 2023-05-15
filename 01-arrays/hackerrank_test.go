package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHourglassSum(t *testing.T) {
	var expected int32 = 28
	var matrix [][]int32
	matrix = [][]int32{
		{-9, -9, -9, 1, 1, 1},
		{0, 9, 0, 4, 3, 2},
		{-9, -9, -9, 1, 2, 3},
		{0, 0, 8, 6, 6, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	actual := hourglassSum(matrix)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
	matrix = [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}
	actual = hourglassSum(matrix)
	expected = -6
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func hourglassSum(arr [][]int32) int32 {
	var max int32
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] +
				arr[i+1][j+1] +
				arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			fmt.Println(sum, max)
			if i == 0 && j == 0 {
				max = sum
			}
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func TestRotLeft(t *testing.T) {
	var expected []int32 = []int32{5, 1, 2, 3, 4}
	var result []int32
	result = rotLeft([]int32{1, 2, 3, 4, 5}, 4)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func rotLeft(a []int32, d int32) []int32 {
	aLen := int32(len(a))
	result := make([]int32, aLen)

	for i := int32(0); i < aLen; i++ {
		newIndex := (i + d) % aLen
		result[i] = a[newIndex]
	}

	return result
}

func TestMinimumBrides(t *testing.T) {
	var expected int32 = 1
	var result int32
	result = minimumBribes([]int32{1, 2, 3, 5, 4, 6, 7, 8})
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	result = minimumBribes([]int32{4, 1, 2, 3})
	expected = -1
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	//2 1 5 3 4
	result = minimumBribes([]int32{2, 1, 5, 3, 4})
	expected = 3
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func minimumBribes(q []int32) int32 {
	var bribes int32 = 0
	var i int32 = 0
	var j int32 = 0
	for i = int32(len(q) - 1); i >= 0; i-- {
		if int32(q[i])-(i+1) > 2 {
			return -1
		}
		for j = max(0, q[i]-2); j < i; j++ {
			if q[j] > q[i] {
				bribes++
			}
		}
	}
	return bribes
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
