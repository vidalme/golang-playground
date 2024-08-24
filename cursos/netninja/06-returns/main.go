package main

import (
	"fmt"
	"strings"
)

func iniciais(nome string) (string, string) {
	a := strings.ToUpper(nome)
	nomes := strings.Split(a, " ")

	var inis []string

	for _, v := range nomes {
		inis = append(inis, v[:1])
	}
	fmt.Printf("array é %q \n", inis)

	if len(inis) > 1 {
		return inis[0], inis[1]
	}
	return inis[0], "_"

}

func main() {
	fn1, sn1 := iniciais("Andre Vidal")
	fn2, sn2 := iniciais("Manu")
	fn3, sn3 := iniciais("Matilda Vidal Mendes")
	fmt.Println(fn1, sn1)
	fmt.Println(fn2, sn2)
	fmt.Println(fn3, sn3)
}
