package main

import "fmt"

var olaPort string = "Olá, "

func Ola(nome string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return olaPort + nome
}

func main() {
	fmt.Println(Ola("andre"))
}
