package main

import (
	"fmt"
)

func areaT(b, a float64) (resultado float64) {
	resultado = (b * a) / 2
	return resultado
}

func main() {
	var base, altura float64

	fmt.Printf("Informe a base do triângulo: ")
	fmt.Scanln(&base)

	fmt.Printf("Informe a altura do triângulo: ")
	fmt.Scanln(&altura)

	res := areaT(base, altura)
	fmt.Println("Área do Triângulo =", res)
}
