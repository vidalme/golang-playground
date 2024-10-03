package interacao

import (
	"testing"
)

func TestRepete(t *testing.T) {
	repetido := Repete("a")
	esperado := "aaaaa"

	if repetido != esperado {
		t.Errorf("Expected %q but got %q", esperado, repetido)
	}

}
