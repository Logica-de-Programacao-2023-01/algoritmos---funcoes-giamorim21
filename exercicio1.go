package main

import "fmt"

//Crie uma função que receba dois parâmetros (a e b)
//e retorne a divisão entre eles. Caso o
//segundo parâmetro seja zero, retorne um erro.

func main() {
	var x, y int
	fmt.Println("Digite um valor para x:")
	fmt.Scan(&x)
	fmt.Println("Digite um valor para y:")
	fmt.Scan(&y)

	resultado, err := divisao(x, y)
	if err != nil {
		fmt.Printf("Ocorreu um erro ao dividir x e y: %s\n", err)
		return
	}
	fmt.Printf(" a divisao dos dois é %d:", resultado)
}

func divisao(x int, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return x / y, nil
}
