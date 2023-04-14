package main

import (
	"fmt"
	"strings"
)

//Crie uma função que receba um slice de strings
//e retorne a concatenação (tirar os espaços)
//de todas as strings.

func main() {
	slice := []string{"Hello", "World", "!"}
	concatenacao := Strings(slice)
	fmt.Printf("O resultado da concatenação de %v é: %s", slice, concatenacao)
}

func Strings(s []string) string {
	return strings.Join(s, "")
}
