package model

import "fmt"

type player struct {
	nome string
	kind string
}

type IPlayer interface {
	TocaMusica(m musica)
	ParaMusica(m musica)
	MudaEstacao(e estacao)
}

func NovoPlayer(n string, k string) player {
	return player{
		nome: n,
		kind: k,
	}
}

func (p player) TocaMusica(m musica) {
	fmt.Println("Toca musica ", m.nome, "do ", m.autor)
}

func (p player) ParaMusica(m musica) {
	fmt.Println("Toca musica ", m.nome, "do ", m.autor)
}

func (p player) MudaEstacao(ne estacao) {
	fmt.Println("Mudando para estação ", ne.nome)
}
