package contas

import (
	"alura/banco/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (t *ContaCorrente) Sacar(q float64) string {
	fmt.Print("helo")
	if t.saldo > q && q > 0 {
		t.saldo -= q
		fmt.Printf("Saque de %v realizado com sucesso.\n", q)
		fmt.Println("saldo atual :", t.saldo)
	} else {
		fmt.Println("Não ha fundos para o saque")
	}
	return "oi"
}

func (t *ContaCorrente) Depositar(q float64) {
	if q > 0 {
		t.saldo += q
	}
	fmt.Printf("Deposito de %v realizado com sucesso.\n", q)
	fmt.Println("saldo atual :", t.saldo)
}

func (t *ContaCorrente) Transferir(a *ContaCorrente, q float64) {
	t.Sacar(q)
	a.Depositar(q)
	fmt.Println("transferencia completa")
}

func (t *ContaCorrente) ObterSaldo() float64 {
	return t.saldo
}
