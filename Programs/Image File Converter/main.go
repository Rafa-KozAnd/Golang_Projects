package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func converterImagem(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, img, nil)
	return err
}

func main() {
	var inputPath, outputPath string
	fmt.Print("Digite o caminho da imagem de entrada: ")
	fmt.Scan(&inputPath)
	fmt.Print("Digite o caminho de saída: ")
	fmt.Scan(&outputPath)

	err := converterImagem(inputPath, outputPath)
	if err != nil {
		fmt.Println("Erro ao converter imagem:", err)
	} else {
		fmt.Println("Imagem convertida com sucesso!")
	}
}
