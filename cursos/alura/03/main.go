package main

import (
	"alura-loja/routes"
	"net/http"

	_ "github.com/lib/pq"
)

<<<<<<< HEAD
func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres password=qwe dbname=alura_loja host=172.23.64.1 sslmode=disable"
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

=======
>>>>>>> refs/remotes/upstream/main
func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8082", nil)
}
