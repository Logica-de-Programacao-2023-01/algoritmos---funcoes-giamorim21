package main

import (
	"errors"
	"fmt"
)

//Escreva uma função que receba um slice de inteiros como parâmetro e
//retorne um novo slice com apenas os números pares contidos no slice.
//Caso o slice esteja vazio, retorne um erro.

func pares(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice está vazia")
	}
	novaSlice := []int{}
	for _, num := range slice {
		if num%2 == 0 {
			novaSlice = append(novaSlice, num)
		}
	}
	return novaSlice, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	novaSlice, err := pares(slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(novaSlice)
	}
}
