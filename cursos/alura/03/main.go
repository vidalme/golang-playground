package main

import (
	"alura-loja/models"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8082", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	var temp = template.Must(template.ParseGlob("templates/*.html"))
	temp.ExecuteTemplate(w, "Index", models.BuscaTodosProdutos())
}
