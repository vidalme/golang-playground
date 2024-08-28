package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres password=12345678 dbname=alura_loja host=localhost sslmode=disable"
	// connStr := "user=postgres password=yourpassword dbname=yourdbname sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error)
	}
	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8082", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	db := conectaComBancoDeDados()

	selectTodosProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(selectTodosProdutos)

	// p := Produto{}
	// produtos := []Produto{}

	// for selectTodosProdutos.Next() {
	// 	var id int
	// 	var nome, descricao string
	// 	var preco float64
	// 	var quantidade int

	// 	err := selectTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	p.Nome = nome
	// 	p.Descricao = descricao
	// 	p.Preco = preco
	// 	p.Quantidade = quantidade

	// 	produtos = append(produtos, p)
	// }

	// if err != nil {
	// 	panic(err.Error())
	// }

	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortavel", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Prancha de surf", "radical", 899, 1},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
