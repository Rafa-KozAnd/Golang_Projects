package main

import "fmt"

func tabuada(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}

func main() {
	var numero int
	fmt.Print("Digite um número para gerar a tabuada: ")
	fmt.Scan(&numero)
	tabuada(numero)
}
