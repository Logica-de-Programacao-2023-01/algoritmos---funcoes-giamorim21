package main

import (
	"errors"
	"fmt"
	"strings"
)

//Escreva uma função que receba um slice de strings como parâmetro e
//retorne uma string contendo apenas as strings que começam com uma
//letra maiúscula. Caso o slice esteja vazio, retorne um erro.

func maiuscula(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice está vazia")
	}
	var result []string
	for _, s := range slice {
		if len(s) > 0 && strings.ToUpper(string(s[0])) == string(s[0]) {
			result = append(result, s)
		}
	}
	return strings.Join(result, " "), nil
}

func main() {
	slice := []string{"Hello", "World", "brasil", "lingua", "DF"}
	a, err := maiuscula(slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}
}
