package array

import "math/big"

func createArray() []int {
	return []int{1, 2, 3, 4, 5}
}

func createArrayFixedSize() [5]int {
	return [5]int{1, 2, 3, 4, 5}
}

// Modificación de un array: Toma el array que creaste en el primer ejercicio y modifica uno de los valores. Luego imprime el array para verificar que el valor ha cambiado.
func modifyArray(numbers []int) [5]int {
	a := [5]int(numbers)
	a[0] = 0
	return a
}

// Recorrido de un array: Usa un bucle para recorrer el array y sumar todos los valores. Imprime el total al final.
func sumArray(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func sumArrayBigInt(numbers []big.Int) *big.Int {
	total := big.NewInt(0)
	for _, n := range numbers {
		total = total.Add(total, &n)
	}
	return total
}

//TODO ejercicos de slices
