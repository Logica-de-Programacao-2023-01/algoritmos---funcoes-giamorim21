package main

import "fmt"

//Crie uma função que receba uma string como parâmetro e retorne o número de
//caracteres contidos nessa string. Caso a string seja vazia,
//retorne um erro.

func main() {
	result, err := palavra("Hello Word")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func palavra(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("string vazia")
	}
	return len(s), nil
}
