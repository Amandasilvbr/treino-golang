//  Modifique o programa anterior para imprimir um triângulo de asteriscos invertido;

package main

import (
	"fmt"
)

func main() {

	num := 0

	fmt.Println("Digite o número de linhas desejado")
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		for j := num; j >= i; j-- {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}