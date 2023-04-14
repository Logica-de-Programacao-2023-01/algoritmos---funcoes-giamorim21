package main

import (
	"errors"
	"fmt"
	"strings"
)

//Escreva uma função que receba um slice de strings
//como parâmetro e retorne uma string com todas as strings
//concatenadas e separadas por vírgulas.
//Caso o slice esteja vazio, retorne um erro.

func main() {
	slice := []string{"maçã", "banana", "laranja"}
	result, err := concatenacaostring(slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func concatenacaostring(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice está vazio")
	}
	return strings.Join(slice, ","), nil
}
