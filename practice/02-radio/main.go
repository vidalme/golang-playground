package main

import (
	"fmt"
	"radio/model"
)

func main() {
	p := model.NovoPlayer("Verdinha", "FM")
	fmt.Print(p)

	p.TocaMusica()
}

func criaMusicas() {
	// musicas := []{}
}

// package main

// import (
// 	"fmt"
// 	"interacao/autores"
// )

// func chamaAutor() {
// 	pn := "Joao"
// 	un := "Soldado"
// 	nac := "Brasileiro"

// 	a := autores.Autor{
// 		PrimeiroNome:  pn,
// 		UltimoNome:    un,
// 		Nacionalidade: nac,
// 	}

// 	fmt.Println(a)
// }

// func main() {
// 	chamaAutor()
// }
