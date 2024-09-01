package models

import (
	"alura-loja/db"
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

	selectTodosProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(selectTodosProdutos)

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

func SelecionaProduto(ide string) Produto {
	db := db.ConectaComBancoDeDados()

	query := "select * from produtos where id=$1"
	selecionaProduto := db.QueryRow(query, ide)

	var id int
	var nome, descricao string
	var preco float64
	var quantidade int

	selecionaProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)

	p := Produto{}
	p.Id = id
	p.Nome = nome
	p.Descricao = descricao
	p.Preco = preco
	p.Quantidade = quantidade

	defer db.Close()
	return p
}

func EditaProduto(id string) Produto {

	db := db.ConectaComBancoDeDados()

	query := "select from produtos where id=$1"
	editaProduto := db.QueryRow(query, id)

	var p Produto
	editaProduto.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)

	return p
}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()

}
