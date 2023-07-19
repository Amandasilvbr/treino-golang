// Escreva uma função que recebe uma string como entrada e retorna uma nova string que é a versão invertida da original.

package main

import "fmt"

func newString(word string) string {
	var finalString string
	for j := len(word) - 1; j >= 0; j-- {
		finalString += string(word[j])
	}
	return finalString
}

func main() {
	var word string

	fmt.Println("Digite a palavra")
	fmt.Scanln(&word)

	result := newString(word)
	fmt.Println(result)
}
