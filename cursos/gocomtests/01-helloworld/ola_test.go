package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s' , esperado '%s'", resultado, esperado)
		}
	}
	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})
	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, esperado, resultado)
	})
	t.Run("em frances", func(t *testing.T) {
		resultado := Ola("Fabien", "frances")
		esperado := "Ça va, Fabien"
		verificaMensagemCorreta(t, esperado, resultado)
	})
	t.Run("em ingles", func(t *testing.T) {
		resultado := Ola("Carl", "ingles")
		esperado := "Sup, Carl"
		verificaMensagemCorreta(t, esperado, resultado)
	})
}
