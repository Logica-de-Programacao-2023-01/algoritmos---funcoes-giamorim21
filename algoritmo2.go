package main

import "fmt"

//Crie uma função que
//receba uma string e retorne a quantidade de vogais presentes.

func main() {
	var s string
	fmt.Println("Digite uma string")
	fmt.Scan(&s)
	vogais := countvogais(s)
	fmt.Printf("A string %s contém %d vogais\n", s, vogais)
}

func countvogais(s string) int {
	count := 0
	for _, char := range s {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}
