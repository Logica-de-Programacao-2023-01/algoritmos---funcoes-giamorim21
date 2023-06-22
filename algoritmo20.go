package main

import (
	"errors"
	"fmt"
)

//Escreva uma função que receba um slice de strings como
//parâmetro e retorne um novo slice contendo apenas as strings
//que possuem mais de 5 caracteres. Caso o slice esteja vazio,
//retorne um erro.

func caracter(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice está vazia")
	}
	resultado := make([]string, 0)
	for _, s := range slice {
		if len(s) > 5 {
			resultado = append(resultado, s)
		}
	}
	return resultado, nil
}

func main() {
	slice := []string{"banana", "laranja", "abacaxi", "maca", "uva", "melancia"}
	f, err := caracter(slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}
}
