package main

import (
	"fmt"
)

func vogais (frase string){
	var vogal,i int =0,0

	for i=0; i<len(frase);i++{
		if frase[i]== 'a' || frase[i]== 'e' || frase[i]== 'i' || frase[i]== 'o' || frase[i]== 'u'{
			vogal++
		}
	}
	fmt.Println("A string possui ",vogal, "vogal(is)!")
}

func main(){
	var  frase string

	fmt.Printf("Digite a string de entrada: ")
	fmt.Scanf("%s",&frase)

	vogais(frase)

}