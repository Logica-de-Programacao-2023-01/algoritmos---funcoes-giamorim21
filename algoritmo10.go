package main

import (
	"errors"
	"fmt"
	"sort"
)

//Escreva uma função que receba um slice de inteiros como parâmetro e
//retorne um novo slice com os valores ordenados de forma crescente
//. Caso o slice esteja vazio, retorne um erro.

func crescente(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice está vazia")
	}
	sort.Ints(slice)
	return slice, nil
}

func main() {
	slice := []int{9, 4, 6, 2, 8, 3, 7, 1, 5}
	ordem, err := crescente(slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ordem)
	}
}
