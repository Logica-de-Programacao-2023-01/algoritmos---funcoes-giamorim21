package main

import "fmt"

//Crie uma função que receba um slice de inteiros
//e um valor inteiro, e retorne a posição do primeiro
//elemento igual ao valor no slice.
//Caso não encontre, retorne -1.

func main() {
	slice := []int{1, 2, 3, 8, 5}
	valor := 8
	index := findIndex(slice, valor)
	if index != -1 {
		fmt.Printf("O valor %d foi encontrado na posição %d do slice\n", valor, index)
	} else {
		fmt.Printf("O valor %d não foi encontrado no slice\n", valor)
	}
}

func findIndex(slice []int, valor int) int {
	for i, num := range slice {
		if num == valor {
			return i
		}
	}
	return -1
}
