package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Andre")
	esperado := "Olá, Andre"

	if resultado != esperado {
		t.Errorf("resultado '%s' , esperado '%s'", resultado, esperado)
	}
}
