package main

import "fmt"

var olaPort string = "Olá, "
var olaEspa string = "Hola, "
var olaFran string = "Ça va, "
var olaIngl string = "Sup, "

func Ola(nome string, lingua string) string {

	if nome == "" {
		nome = "Mundo"
	}
	return checkLingua(lingua) + nome
}

func checkLingua(lingua string) string {
	switch lingua {
	case "espanhol":
		return olaEspa
	case "frances":
		return olaFran
	case "ingles":
		return olaIngl
	default:
		return olaPort
	}

}

func main() {
	fmt.Println(Ola("", ""))
	fmt.Println(Ola("Andre", ""))
	fmt.Println(Ola("Elodie", "espanhol"))
	fmt.Println(Ola("Andre", "frances"))
	fmt.Println(Ola("Bill", "ingles"))
}
