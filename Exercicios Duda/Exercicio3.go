package main

import (
	"errors"
	"fmt"
)

func compara(a int) (string, error) {
	if a < 0 {
		err := errors.New("Valor Negativo!")
		return "", err
	} else {
		return "Valor Positivo!", nil
	}
}

func main() {
	var numero int

	fmt.Println("Digite um número inteiro: ")
	fmt.Scanf("%d", &numero)

	resultado, err := compara(numero)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println(resultado)
	}
}
