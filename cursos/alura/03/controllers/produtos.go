package controllers

import (
	"alura-loja/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Index", models.BuscaTodosProdutos())
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(nome, descricao, preco, quantidade)
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preço:", err)
		}
		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao do quantidade:", err)
		}
		models.CriaNovoProduto(nome, descricao, precoFloat, quantidadeInt)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")
	tmpl.ExecuteTemplate(w, "Edit", models.SelecionaProduto(idProduto))
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversao do ID para Int: ", err)
		}
		precoConvFlo, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do Preco para Float64: ", err)
		}
		quantConvInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao do Quantidade para Int: ", err)
		}
		models.AtualizaProduto(idConvInt, nome, descricao, precoConvFlo, quantConvInt)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
