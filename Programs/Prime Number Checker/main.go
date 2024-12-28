package main

import "fmt"

func isPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var numero int
	fmt.Print("Digite um número para verificar se é primo: ")
	fmt.Scan(&numero)

	if isPrimo(numero) {
		fmt.Println("O número é primo!")
	} else {
		fmt.Println("O número não é primo.")
	}
}
