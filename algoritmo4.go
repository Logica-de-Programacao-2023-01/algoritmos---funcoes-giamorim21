package main

import (
	"fmt"
	"sort"
)

//Crie uma função que receba um slice de inteiros
//e retorne o segundo maior valor.

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:")
	fmt.Scan(&slice)
	segundoMaior := segundoMaior(slice)
	fmt.Println("Segundo maior valor:", segundoMaior)
}

func segundoMaior(slice []int) int {
	if len(slice) < 2 {
		panic("Slice deve ter pelo menos dois elementos")
	}
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	return slice[1]
}
