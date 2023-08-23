package fun_facts

import (
	"fmt"
)

func deferer() {
	for i := 0; i < 5; i++ {
		defer println(i) // 4 3 2 1 0
	}
}

func copier() {
	data := []int{1, 2, 3, 4, 5, 6}
	mapa := make([]int, 0) // if you create it with 5 size, it create an array filled by 0. Change 5 -> 0 or use var
	for _, d := range data {
		mapa = append(mapa, d)
	}
	fmt.Println(mapa)
}

func adder() {
	var num int8 = 127 // binary representation: 01111111
	res := num + 1     // when add 1, the new representation is: 10000000
	fmt.Println(res)   // this binary nubmer is -128, the first bit
}

func slicer() {
	data := []int{1, 2, 3, 4, 5, 6}
	res := data[:3]
	fmt.Println(res, len(res), cap(res))
	final := res[2:4] //making slice from slice doesn't make a copy
	fmt.Println(final, len(final), cap(final))
}
