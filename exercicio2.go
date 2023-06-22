package main

import "fmt"

//Escreva uma função que calcule a média de uma lista de números
//e retorne um erro caso
//a lista esteja vazia.

func main() {
	result, err := media([]int{1, 2, 3, 4, 5})
	if err != nil {
		fmt.Printf("Houve um erro: %s\n", err)
	} else {
		fmt.Printf("A media é de %d", result)
	}

}

func media(list []int) (int, error) {
	if len(list) == 0 {
		return 0, fmt.Errorf("A lista está vazia")
	}
	soma := 0
	for _, x := range list {
		soma += x
	}
	return soma / len(list), nil
}