package models

import (
	"alura-loja/db"
	"fmt"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectTodosProdutos, err := db.Query("select * from Produtos")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(selectTodosProdutos)

	p := Produto{}
	produtos := []Produto{}

	for selectTodosProdutos.Next() {
		var id int
		var nome, descricao string
		var preco float64
		var quantidade int

		err := selectTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}
