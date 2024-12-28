package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Jogador struct {
	nome   string
	vida   int
	ataque int
}

type Inimigo struct {
	nome   string
	vida   int
	ataque int
}

func (j *Jogador) atacar(i *Inimigo) {
	dano := rand.Intn(j.ataque)
	i.vida -= dano
	fmt.Printf("%s atacou %s causando %d de dano!\n", j.nome, i.nome, dano)
}

func (i *Inimigo) atacar(j *Jogador) {
	dano := rand.Intn(i.ataque)
	j.vida -= dano
	fmt.Printf("%s atacou %s causando %d de dano!\n", i.nome, j.nome, dano)
}

func batalha(j *Jogador, i *Inimigo) {
	for j.vida > 0 && i.vida > 0 {
		j.atacar(i)
		if i.vida <= 0 {
			fmt.Printf("%s derrotou o inimigo %s!\n", j.nome, i.nome)
			return
		}
		i.atacar(j)
		if j.vida <= 0 {
			fmt.Printf("%s foi derrotado pelo inimigo %s...\n", j.nome, i.nome)
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	jogador := Jogador{
		nome:   "Herói",
		vida:   100,
		ataque: 20,
	}

	inimigo := Inimigo{
		nome:   "Goblin",
		vida:   50,
		ataque: 15,
	}

	fmt.Println("Bem-vindo ao Dungeon Crawl!")
	batalha(&jogador, &inimigo)
}
