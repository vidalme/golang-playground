package main

import (
	"fmt"
	"sort"
)

//teste

func main() {
	// mensagem := "oi terraqueos"
	// fmt.Println(mensagem)
	// fmt.Println(strings.Contains(mensagem, "oi"))
	// fmt.Println(strings.ReplaceAll(mensagem, "oi", "olá"))
	// fmt.Println(mensagem)

	// fmt.Println(strings.ToUpper(mensagem))
	// fmt.Println(strings.Split(mensagem, " "))

	numeros := []int{100, 20, 30, 44, 22, 2000, 774, 23}

	sort.Ints(numeros)
	fmt.Println(numeros)

}
