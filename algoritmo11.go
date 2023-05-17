package main

import (
	"errors"
	"fmt"
	"math"
)

//Crie uma função que receba um número inteiro como parâmetro e
//retorne verdadeiro se ele for um número primo e falso caso contrário.
//Caso o número seja negativo, retorne um erro.

func primo(x int) (bool, error) {
	if x < 0 {
		return false, errors.New("O número é negativo")
	}
	if x < 2 {
		return false, nil
	}
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func main() {
	x := 5
	numero, err := primo(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numero)
	}
}
