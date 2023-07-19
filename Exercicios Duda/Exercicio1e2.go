package main

import (
	"fmt"
)

//Exercicio 1
func FazTriangulo (a int){
	var i,j,b int = 0,0,0;
	fmt.Printf("Triangulo Normal:\n");
	for i=0;i<a;i++{
		b++;
		for j=0;j<b;j++{
			fmt.Print("*");
		}
		fmt.Print("\n");
	}
}

//Exercicio 2
func FazTrianguloInvertido (a int){
	fmt.Printf("\nTriângulo Invertido: \n");
	var i,j,b int = 0,0,a;
	for i=0;i<a;i++{
		for j=0;j<b;j++{
			fmt.Print("*");
		}
		b--;
		fmt.Print("\n");
	}
}

func main (){
	var numero int = 0
fmt.Println("Digite o número de linhas desejados: ");
fmt.Scanf("%d",&numero);

//Exercicio 1
FazTriangulo(numero);

//Exercicio 2
FazTrianguloInvertido(numero);

}