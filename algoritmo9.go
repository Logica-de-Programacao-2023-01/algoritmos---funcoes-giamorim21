package main

import (
	"errors"
	"fmt"
	"strings"
)

//Crie uma função que receba uma string como parâmetro e retorne um
//novo slice com todas as palavras contidas na string. Considere que
//as palavras são separadas por espaços em branco. Caso a string seja
//vazia, retorne um erro.

func separarPalavras(s string) ([]string, error) {
	if len(s) == 0 {
		return nil, errors.New("A string está vazia")
	}
	palavras := strings.Split(s, " ")
	return palavras, nil
}

func main() {
	s := "Olá, tube bem?"
	palavra, err := separarPalavras(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(palavra)
	}
}
