// package main

// import "testing"

// func TestSoma(t *testing.T) {

// 	//act
// 	want := 3 + 3

// 	//arrange
// 	got := Soma(3, 3)

// 	//assert
// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

package main

import "testing"

func TestSoma(t *testing.T) {
	t.Run("soma simples", func(t *testing.T) {

		got := Soma(3, 3)
		want := 3 + 3

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
	t.Run("soma de numero negativo", func(t *testing.T) {

		got := Soma(-1, 3)
		want := -1 + 3

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}
