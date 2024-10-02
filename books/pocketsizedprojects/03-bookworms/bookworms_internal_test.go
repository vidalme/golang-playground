package main

import "testing"

var (
	gatoviajante = Book{Author: "Hiro Arikawa", Title: "Relatos de um gato viajante"}
	pum          = Book{Author: "Blandina Franco", Title: "Quem soltou o pum?"}
	ernesto      = Book{Author: "Blandina Franco", Title: "Ernesto"}
	dune         = Book{Author: "Frank Herbert", Title: "Dune"}
	republica    = Book{Author: "Platão", Title: "A república"}
)

func TestLoadBoookworms(t *testing.T) {

	type testCase struct {
		bookormsFile string
		want         []Bookworm
		wantErr      bool
	}

	tests := map[string]testCase{

		"file exists": {
			wantErr:       false,
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Manu", Books: []Book{gatoviajante}},
				{Name: "Matilda", Books: []Book{pum, ernesto}},
				{Name: "Andre", Books: []Book{dune, republica}},
			},
		},
		"file does not exit": {
			wantErr:       true,
			bookwormsFile: "testdata/sem_arquivo_aqui.json",
			want:          nil,
		},
		"json invalido": {
			wantErr:      true,
			bookormsFile: "testdata/invalido.json",
			want:         nil,
		},
	}
}
