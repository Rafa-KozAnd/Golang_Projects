package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Ação struct {
	nome   string
	preço  float64
}

func gerarPrecoAleatorio() float64 {
	return rand.Float64() * 100 // Preço entre 0 e 100
}

func simularMercado() []Ação {
	ações := []string{"Ação A", "Ação B", "Ação C", "Ação D", "Ação E"}
	mercado := []Ação{}

	for _, nome := range ações {
		mercado = append(mercado, Ação{
			nome:  nome,
			preço: gerarPrecoAleatorio(),
		})
	}
	return mercado
}

func main() {
	rand.Seed(time.Now().UnixNano())
	mercado := simularMercado()

	fmt.Println("Bem-vindo ao simulador de Mercado de Ações!")
	fmt.Println("Ações disponíveis:")
	for _, acao := range mercado {
		fmt.Printf("%s: R$%.2f\n", acao.nome, acao.preço)
	}
}
