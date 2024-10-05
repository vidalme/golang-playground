package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

type movie struct {
	Nome    string `json:"nome"`
	Diretor string `json:"diretor"`
	Duracao int    `json:"duracao"`
}

type sala struct {
	nome    string
	filme   movie
	sessoes int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var fp string = "dados/filmes_atuais.json"

func coletaFilmes() []movie {

	o, e := os.ReadFile(fp)
	check(e)

	filmes := []movie{}
	json.Unmarshal(o, &filmes)

	return filmes
}

func selecionaFilmeAleatoriamente(ml []movie) movie {
	rn := rand.Intn(len(ml))
	return ml[rn]
}

func passaFilme(m movie, ch chan string) {
	short_version := m.Duracao / 30
	d := time.Second * time.Duration(short_version)
	fmt.Printf("Começou na sala -> Filme %s do diretor %s com duração %d segundos.\n", m.Nome, m.Diretor, short_version)
	time.Sleep(d)
	fmt.Printf("Terminou na sala %s ->  Filme %s do diretor %s com duração %d segundos.\n", <-ch, m.Nome, m.Diretor, short_version)
}

func criaSalas(ml []movie) []sala {

	names := []string{"A", "B", "C"}
	salas := make([]sala, 3)

	for i, v := range names {
		salas[i] = sala{
			nome:    v,
			filme:   selecionaFilmeAleatoriamente(ml),
			sessoes: 0,
		}
	}
	return salas
}

func fechaCinema(ml *[]movie) {
	var md float64
	for _, v := range *ml {
		md = math.Max(float64(v.Duracao), md)
	}
	t := 1 //int(md) / 30
	time.Sleep(time.Second * time.Duration(t))
	fmt.Printf("Cinema fechou após %d segundos\n", t)
}

func abreCinema(salas []sala) {

	mc := make(chan string, 3)

	for _, sala := range salas {
		mc <- sala.nome
		go passaFilme(sala.filme, mc)
	}

}

func main() {
	listaFilmes := coletaFilmes()
	salas := criaSalas(listaFilmes)

	abreCinema(salas)

	fechaCinema(&listaFilmes)
}
