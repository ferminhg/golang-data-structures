package array

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/bxcodec/faker"
)

func TestCreateArrayFixedSize(t *testing.T) {
	expected := [5]int{1, 2, 3, 4, 5}
	result := createArrayFixedSize()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
	t.Log("create array success")
}

func TestCreateArray(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	result := createArray()

	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestModifyArray(t *testing.T) {
	expected := [5]int{0, 2, 3, 4, 5}
	result := modifyArray(createArray())
	if !reflect.DeepEqual(expected, result) {
		t.Error("Expected ", expected, ", but got ", result)
	}
}

func TestSumArray(t *testing.T) {
	expected := 15
	result := sumArray(createArray())
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSumBigArray(t *testing.T) {
	expected := big.NewInt(15)
	numbers := []big.Int{*big.NewInt(15)}
	result := sumArrayBigInt(numbers)

	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	faker := BigIntArrayFacker(5)
	if len(faker) != 5 {
		t.Errorf("Expected %v, but got %v", 5, len(faker))
	}

	result = sumArrayBigInt(faker)

	if result.Cmp(&faker[0]) > 1 {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func BigIntArrayFacker(size int) []big.Int {
	var fakeNumbers []big.Int
	for i := 0; i < size; i++ {
		var num int
		faker.FakeData(&num)
		fakeNumbers = append(fakeNumbers, *big.NewInt(int64(num)))
	}
	return fakeNumbers
}

func TestNameStringArray(t *testing.T) {
	names := createNameArray()
	for _, name := range names {
		t.Log(name)
	}

	if len(names) != 3 {
		t.Errorf("Expected %v, got %v", 3, len(names))
	}
}

func createNameArray() []string {
	return []string{"Lucia", "Maya", "Beka"}
}

func TestSumMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := 45
	result := sumMatrix(&matrix)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func sumMatrix(matrix *[][]int) int {
	sum := 0
	for _, row := range *matrix {
		for _, num := range row {
			sum += num
		}
	}
	return sum
}

func TestSearchOnArray(t *testing.T) {
	emptyArray := []int{}
	var result bool
	result = search(0, emptyArray)
	if result != false {
		t.Errorf("Expected %v, got %v", false, result)
	}

	fixedArray := createArrayFixedSize()
	result = search(1, fixedArray[:])
	if result != true {
		t.Errorf("Expected %v, got %v", true, result)
	}
}

func search(number int, array []int) bool {
	for _, num := range array {
		if num == number {
			return true
		}
	}
	return false
}
