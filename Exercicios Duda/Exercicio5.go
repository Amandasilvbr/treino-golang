package main

import (
	"fmt"
)

func fatorial(a int) {
	i := 1
	resultado := 1
	for i = 1; i <= a; i++ {
		resultado *= i
	}

	fmt.Println("Fatorial de ", a, "=", resultado)
}

func main() {
	numero := 0
	fmt.Printf("Digite um valor: \n")
	fmt.Scanf("%d", &numero)

	fatorial(numero)

}
