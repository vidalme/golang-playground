package main

import (
	"bufio"
	"fmt"
	"golang-playground/cursos/alura/02/contas"
	"os"
	"strings"
)

func main() {

	contaAndre := contas.ContaBancaria{Titular: "Andre", Saldo: 200}
	contaManu := contas.ContaBancaria{Titular: "Manu", Saldo: 300}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("aperte - 1 saque , 2 deposito, 3 tranferencia, 0 sair")
		a, _ := reader.ReadString('\n')
		a = strings.TrimSpace(a)
		fmt.Println(a)

		switch a {
		case "1":
			fmt.Println("escolheu 1")
			contaAndre.saque(100)
		case "2":
			fmt.Println("escolheu 2")
			contaAndre.depositar(100)
		case "3":
			fmt.Println("escolheu 3")
			contaAndre.transferir(&contaManu, 50)
		case "0":
			os.Exit(1)
		default:
			fmt.Println("default")
		}

	}

}
