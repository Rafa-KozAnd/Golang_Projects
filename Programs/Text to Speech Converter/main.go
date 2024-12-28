package main

import (
	"fmt"
	"os/exec"
)

func textoParaFala(texto string) {
	cmd := exec.Command("espeak", texto)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao executar comando:", err)
	}
}

func main() {
	var texto string
	fmt.Print("Digite o texto para conversão em fala: ")
	fmt.Scan(&texto)

	textoParaFala(texto)
}
