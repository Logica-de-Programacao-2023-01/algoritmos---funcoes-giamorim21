package main

import "fmt"

//Escreva uma função que receba um
//slice de inteiros como parâmetro
// e retorne o menor
//valor contido no slice.
//Caso o slice esteja vazio, retorne um erro.

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	menor, err := lista(slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(menor)
	}
}

func lista(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, fmt.Errorf("slice vazio")
	}
	menor := slice[0]
	for _, valor := range slice {
		if valor < menor {
			menor = valor
		}
	}
	return menor, nil
}
