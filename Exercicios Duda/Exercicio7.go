package main

import (
	"fmt"
)

func inverte(frase string) (novas string) {

	for i := len(frase) - 1; i >= 0; i-- {
		novas += string(frase[i])
	}
	return novas
}

func main() {
	var frase, frasenova string

	fmt.Printf("Digite a string de entrada: ")
	fmt.Scanf("%s", &frase)

	frasenova = inverte(frase)
	fmt.Println(frasenova)
}
