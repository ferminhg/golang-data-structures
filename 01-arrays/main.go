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

// Array de strings: Crea un array de strings y úsalo para almacenar nombres de personas. Luego recorre el array e imprime cada nombre en una línea separada.

// Array de arrays (Matrices): Crea una matriz (un array de arrays) de enteros. Puedes pensar en esto como una tabla de números con filas y columnas. Intenta sumar todos los números en la matriz.

// Búsqueda en un array: Crea un array de 10 números. Luego, escribe una función que tome un número y devuelva true si el número está en el array y false si n

//TODO ejercicos de slices
