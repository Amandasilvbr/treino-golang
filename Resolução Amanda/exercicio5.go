// Escreva uma função que recebe um número inteiro positivo como entrada e calcula o fatorial desse número. O fatorial de um número n é o produto de todos os números inteiros de 1 a n.

package main

import (
	"errors"
	"fmt"
)

func fatorial(num int) (int, error) {
	resultado := 1

	err := errors.New("Número negativo")

	if num < 0 {
		return 0, err

	} else if num == 0 {
		return 1, nil

	} else {
		for i := 0; i < num; i++ {
			resultado += resultado * i
		}
		return resultado, nil
	}
}

func main() {
	var num int

	fmt.Println("Digite um número inteiro positivo")
	fmt.Scanln(&num)

	fatorial, err := fatorial(num)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("O resultado do fatorial é %d", fatorial)
	}
}
