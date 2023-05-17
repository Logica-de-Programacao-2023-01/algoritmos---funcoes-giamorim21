package main

import (
	"errors"
	"fmt"
	"strings"
)

//Crie uma função que receba um slice de strings como parâmetro e
//retorne uma nova string com as strings em ordem alfabética. Caso
//o slice esteja vazio, retorne um erro.

func alfabetica(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("A slice está vazia")
	}
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if slice[j] > slice[i] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	resultado := strings.Join(slice, " ")
	return resultado, nil
}

func main() {
	slice := []string{"banana", "laranja", "abacaxi", "maçã"}
	c, err := alfabetica(slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
}
