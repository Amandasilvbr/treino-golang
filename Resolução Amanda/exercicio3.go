//  Crie uma função personalizada que recebe um número inteiro positivo como entrada. A função deve retornar o resultado de uma operação qualquer realizada com esse número. No entanto, se o número de entrada for negativo, a função deve retornar um erro indicando que um número negativo foi fornecido.

package main

import (
	"errors"
	"fmt"
)

func positive(num int) (string, error) {
	if num < 0 {
		err := errors.New("Número negativo")
		return "", err
	} else {
		return "Número positivo", nil
	}
}

func main() {

	num := 0

	fmt.Println("Digite um número inteiro positivo")
	fmt.Scanln(&num)

	result, err := positive(num)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println(result)
	}
}
