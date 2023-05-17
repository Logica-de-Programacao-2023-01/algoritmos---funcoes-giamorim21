package main

import (
	"errors"
	"fmt"
)

//Crie uma função que receba um número inteiro como parâmetro e retorne
//a soma de todos os seus dígitos. Caso o número seja negativo, retorne
//um erro.

func soma(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Número negativo")
	}
	var soma = 0
	for n != 0 {
		digito := n % 10
		soma += digito
		n /= 10
	}
	return soma, nil
}

func main() {
	n := 63
	s, err := soma(n)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
