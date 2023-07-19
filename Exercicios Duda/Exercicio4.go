package main

import (
	"fmt"
)

func main() {
	pessoa := make(map[string]int)
		pessoa["Ana"]     = 20
		pessoa["Pedro"]   = 55
		pessoa["Beatriz"] = 11
		pessoa["Marcio"]  = 60

	for nome,idade:= range pessoa{
		fmt.Println("\nNome: ",nome,"\nIdade: ", idade)
	}

}