package main

import (
	"fmt"
)

func verificaNum(num int) string {
	if num > 0 {
		return "Positivo!"
	}
	if num < 0 {
		return "Negativo!"
	} else {
		return "Zero!"
	}
}

func main() {
	var numero int

	fmt.Printf("Informe um número: ")
	fmt.Scanln(&numero)

	fmt.Println("O número é", verificaNum(numero))
}
