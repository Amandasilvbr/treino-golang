// Escreva uma função que recebe uma lista de números inteiros como entrada e retorna o valor máximo e mínimo encontrados nessa lista.

package main

import "fmt"

func num(numbers []int) ([]int,[]int) {
	max := numbers[0]
	min := numbers[0]

	for _, number := range numbers {
		if number > max {
			max = number
		} else if number < min {
			min = number
		}
	}
	return []int{max}, []int{min}
}

func main() {
	var numbers []int

	fmt.Println("Digite os números")
	fmt.Scanln(&numbers)

	result1, result2 := num(numbers)
	fmt.Println(result1, result2)
}