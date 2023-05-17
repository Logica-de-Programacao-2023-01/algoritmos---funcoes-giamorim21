package main

import (
	"errors"
	"fmt"
)

//Escreva uma função que receba dois slices de inteiros como parâmetros
//e retorne um novo slice contendo apenas os números presentes nos dois
//slices. Caso um dos slices esteja vazio, retorne um erro.

func newSlice(slice1 []int, slice2 []int) ([]int, error) {
	if len(slice1) == 0 && len(slice2) == 0 {
		return nil, errors.New("Uma das slices está vazia")
	}
	numMap := make(map[int]bool)
	for _, num := range slice1 {
		numMap[num] = true
	}
	var resultado []int
	for _, num := range slice2 {
		if numMap[num] {
			resultado = append(resultado, num)
		}
	}
	return resultado, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}
	i, err := newSlice(slice1, slice2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}
