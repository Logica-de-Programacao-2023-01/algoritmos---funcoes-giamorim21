package main

import (
	"errors"
	"fmt"
)

//Crie uma função que receba um número inteiro e um slice de inteiros
//como parâmetros e retorne verdadeiro se o número estiver presente
//no slice e falso caso contrário. Caso o slice esteja vazio, retorne
//um erro.

func parametro(x int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("Slice está vazia")
	}
	for _, num := range slice {
		if num == x {
			return true, nil
		}
	}
	return false, nil
}

func main() {
	x := 5
	slice := []int{1, 2, 3, 4, 5, 6}
	c, err := parametro(x, slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
}
