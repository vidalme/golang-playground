package main

import (
	"alura/banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	// clienteMatilda := clientes.Titular{
	// 	Nome:      "Matilda Vidal",
	// 	CPF:       "123.456.789.99",
	// 	Profissao: "Desenvolvedora",
	// }

	// contaMatilda := contas.ContaCorrente{
	// 	Titular:       clienteMatilda,
	// 	NumeroAgencia: 123,
	// 	NumeroConta:   123456,
	// 	// saldo:         1000,
	// }
	contaMatilda := contas.ContaCorrente{}
	contaBruna := contas.ContaPoupanca{}

	contaMatilda.Depositar(1000)
	contaBruna.Depositar(400)

	fmt.Println(contaMatilda.ObterSaldo())
	fmt.Println(contaBruna.ObterSaldo())

	PagarBoleto(&contaMatilda, 155)
	PagarBoleto(&contaBruna, 60)

	fmt.Println(contaMatilda.ObterSaldo())
	fmt.Println(contaBruna.ObterSaldo())
	// contaAndre := contas.ContaCorrente{Titular: "Andre", Saldo: 200}
	// contaManu := contas.ContaCorrente{Titular: "Manu", Saldo: 300}

	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	fmt.Println("aperte - 1 saque , 2 deposito, 3 tranferencia, 0 sair")
	// 	a, _ := reader.ReadString('\n')
	// 	a = strings.TrimSpace(a)
	// 	fmt.Println(a)

	// 	switch a {
	// 	case "1":
	// 		fmt.Println("escolheu 1")
	// 		contaAndre.Saque(100)
	// 	case "2":
	// 		fmt.Println("escolheu 2")
	// 		contaAndre.Depositar(100)
	// 	case "3":
	// 		fmt.Println("escolheu 3")
	// 		contaAndre.Transferir(&contaManu, 50)
	// 	case "0":
	// 		os.Exit(1)
	// 	default:
	// 		fmt.Println("default")
	// 	}

	// }

}
