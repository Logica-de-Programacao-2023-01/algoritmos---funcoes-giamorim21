package main

import (
	"errors"
	"fmt"
)

//Crie uma função que receba um slice de inteiros
//e uma função como parâmetros.
//A função deve aplicar a função recebida em cada
//elemento do slice e retornar um novo slice
//com os resultados. Caso o slice esteja vazio, retorne um erro.

func main() {
	slice := []int{1, 2, 3, 4, 5}
	result, err := funcao(slice, double)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
func funcao(slice []int, f func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice está vazia")
	}
	result := make([]int, len(slice))
	for i, num := range slice {
		result[i] = f(num)
	}
	return result, nil
}

func double(x int) int {
	return x * 2
}
