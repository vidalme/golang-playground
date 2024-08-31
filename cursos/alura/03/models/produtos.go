package models

import (
	"alura-loja/db"
	"fmt"
)

type Produto struct {
	Id         int
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
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func CriaNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	// fmt.Print(nome, descricao, preco, quantidade)

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}
func DeletaProduto(id string) {

	db := db.ConectaComBancoDeDados()
	deletarProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarProduto.Exec(id)
	defer db.Close()
}
