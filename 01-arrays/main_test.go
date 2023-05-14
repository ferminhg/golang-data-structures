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
