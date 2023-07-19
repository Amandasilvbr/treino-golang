package main

import (
	"fmt"
)

func maxmin(numeros []int) (max, min int) {
	var i int = 1
	max = numeros[0]
	min = numeros[0]
	for i = 1; i < len(numeros); i++ {
		if numeros[i] > max {
			max = numeros[i]
		}
		if numeros[i] < min {
			min = numeros[i]
		}
	}
	return max, min
}

func main() {
	var numeros []int
	var r1, r2, quantidade, valor, i int

	fmt.Println("Informe quantos valores serão informados: ")
	fmt.Scanln(&quantidade)

	for i = 0; i < quantidade; i++ {
		fmt.Println("Informe os valores desejados: ")
		fmt.Scanln(&valor)
		numeros = append(numeros, valor)
	}

	r1, r2 = maxmin(numeros)
	fmt.Println("Maior valor:", r1, "\nMenor valor:", r2)
}
