// Escreva um programa em Go que utilize uma estrutura de dados map para armazenar nomes de pessoas como chaves e suas idades como valores. Imprima o mapa completo.

package main

import (
	"fmt"
)

func main() {
	people := make(map[string]int)

	people["a"] = 50
	people["aa"] = 40
	people["aaa"] = 30

	for name, age := range people {
		fmt.Printf("A idade da pessoa %s é %d\n", name, age)
	}

}
