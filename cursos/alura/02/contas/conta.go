package contas

import "fmt"

type ContaBancaria struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (t *ContaBancaria) saque(q float64) {
	fmt.Print("helo")
	if t.Saldo > q && q > 0 {
		t.Saldo -= q
		fmt.Printf("Saque de %v realizado com sucesso.\n", q)
		fmt.Println("Saldo atual :", t.Saldo)
	} else {
		fmt.Println("Não ha fundos para o saque")
	}
}

func (t *ContaBancaria) depositar(q float64) {
	if q > 0 {
		t.Saldo += q
	}
	fmt.Printf("Deposito de %v realizado com sucesso.\n", q)
	fmt.Println("Saldo atual :", t.Saldo)
}

func (t *ContaBancaria) transferir(a *ContaBancaria, q float64) {
	t.saque(q)
	a.depositar(q)
	fmt.Println("transferencia completa")
}
