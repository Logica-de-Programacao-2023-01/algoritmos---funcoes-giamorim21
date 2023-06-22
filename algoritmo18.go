package main

import (
	"errors"
	"fmt"
)

//Escreva uma função que receba um slice de inteiros e uma
//função como parâmetros. A função deve aplicar a função
//recebida em cada elemento do slice e retornar a soma de
//todos os resultados. Caso o slice esteja vazio, retorne um
//erro.

func Soma(slice []int, f func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("Slice está vazia")
	}
	soma := 0
	for _, num := range slice {
		soma += f(num)
	}
	return soma, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	multiplicar := func(num int) int {
		return num * 3
	}
	resultado, err := Soma(slice, multiplicar)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
