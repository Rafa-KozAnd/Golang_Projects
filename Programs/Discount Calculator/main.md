package main

import "fmt"

func calcularDesconto(preco, desconto float64) float64 {
	return preco - (preco * desconto / 100)
}

func main() {
	var preco, desconto float64
	fmt.Print("Digite o preço do produto: ")
	fmt.Scan(&preco)
	fmt.Print("Digite o percentual de desconto: ")
	fmt.Scan(&desconto)

	precoFinal := calcularDesconto(preco, desconto)
	fmt.Printf("Preço final com desconto: R$%.2f\n", precoFinal)
}
