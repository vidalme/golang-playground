package main

import (
	"fmt"
	"hello/hello"
)

func main() {
	f := hello.Familia{
		Sobrenome: "Vidal",
		Membros:   []string{"Andre", "Manu", "Matilda"},
	}
	fmt.Println(hello.Say(f))
}
