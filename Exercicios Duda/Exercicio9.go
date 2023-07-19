package main

import (
	"fmt"
)

func areaT(b, a int) (resultado int) {
	resultado = (b * a) / 2
	return resultado
}

func main() {
	var base, altura, res int

	fmt.Printf("Informe a base do triângulo: ")
	fmt.Scanln(&base)

	fmt.Printf("Informe a altura do triângulo: ")
	fmt.Scanln(&altura)

	res = areaT(base, altura)
	fmt.Println("Área do Triângulo =", res)
}
