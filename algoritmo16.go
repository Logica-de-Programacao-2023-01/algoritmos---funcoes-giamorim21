package main

import (
	"errors"
	"fmt"
	"strings"
)

//Escreva uma função que receba uma string como parâmetro e retorne
//uma nova string com todas as vogais substituídas por "*". Caso
//a string seja vazia, retorne um erro.

func Vogais(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("A String está vazia")
	}
	vogais := "aeiouAEIOU"
	for _, char := range vogais {
		s = strings.ReplaceAll(s, string(char), "*")
	}
	return s, nil
}

func main() {
	s := "banana"
	vogal, err := Vogais(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(vogal)
	}
}
