package main

import (
	"errors"
	"fmt"
)

//Crie uma função que receba um número inteiro como parâmetro
//e retorne um novo slice contendo todos os números primos
//menores ou iguais a ele. Caso o número seja negativo,
//retorne um erro.

func isPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func lorota(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("Número é negativo")
	}
	primos := make([]int, 0)
	for i := 2; i <= n; i++ {
		if isPrimo(i) {
			primos = append(primos, i)
		}
	}
	return primos, nil
}

func main() {
	n := 27
	p, err := lorota(n)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p)
	}
}
