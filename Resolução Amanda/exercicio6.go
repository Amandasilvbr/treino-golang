//Escreva uma função que recebe uma string como entrada e conta o número de vogais presentes nessa string. Considere vogais apenas as letras "a", "e", "i", "o" e "u".

package main

import "fmt"


func count(word string) int{
	var num int

	for _, letter := range word {
		if letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' {
			num ++
		}
	}
	return num
}

func main() {

	var word string 

	fmt.Println("Digite a palavra")
	fmt.Scanln(&word)

	result := count(word)
	fmt.Println(result)
}