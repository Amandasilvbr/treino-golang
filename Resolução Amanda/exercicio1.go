// Escreva um programa em Go que imprima um triângulo de asteriscos de acordo com o número de linhas que o usuário desejar

package main

import (
	"fmt"
)

func main() {

	num := 0

	fmt.Println("Digite o número de linhas desejado")
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}
