package main

import (
	"fmt"
	"time"
)

func main() {
	var dataString string
	fmt.Print("Digite uma data no formato YYYY-MM-DD: ")
	fmt.Scan(&dataString)

	data, err := time.Parse("2006-01-02", dataString)
	if err != nil {
		fmt.Println("Erro ao parse da data:", err)
		return
	}

	agora := time.Now()
	diff := agora.Sub(data)

	fmt.Printf("Tempo desde a data fornecida: %v\n", diff)
}
