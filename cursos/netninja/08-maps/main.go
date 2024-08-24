package main

import "fmt"

func main() {

	menu := map[string]float64{
		"salada":    5.99,
		"peixe":     13.99,
		"camarão":   23.99,
		"picanha":   15.99,
		"suco":      0.99,
		"sobremesa": 2.29,
	}

	for k, v := range menu {
		fmt.Println(k, " - ", v)
	}

	fones := map[int]string{
		9998811: "Andre",
		9998822: "Manu",
		9998833: "Matilda",
	}

	for k, v := range fones {
		fmt.Println(k, " - ", v)
	}
	fmt.Println("------------")
	fmt.Println(fones[9998811])

	fones[2] = "novo"

	for k, v := range fones {
		fmt.Println(k, " - ", v)
	}

	fones[9998811] = "Andre Vidal"

	for k, v := range fones {
		fmt.Println(k, " - ", v)
	}

	fmt.Println("------------")
	for k, v := range planetas {
		fmt.Printf("planetal %v tem o apelido %v \n", v, k)
	}

}
