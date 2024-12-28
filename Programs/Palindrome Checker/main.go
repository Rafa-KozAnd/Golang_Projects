package main

import "fmt"

func isPalindromo(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var palavra string
	fmt.Print("Digite uma palavra para verificar se é um palíndromo: ")
	fmt.Scan(&palavra)

	if isPalindromo(palavra) {
		fmt.Println("É um palíndromo!")
	} else {
		fmt.Println("Não é um palíndromo.")
	}
}
