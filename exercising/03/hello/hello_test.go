package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Print(t)

	subtestes := []struct {
		description string
		got         string
		want        string
		validate    func(f Familia) bool
	}{
		{
			description: "Checa formatação da mensagem final",
			got: Say(Familia{
				Sobrenome: "Vidal Mendes",
				Membros:   []string{"Andre", "Manu", "Matilda"},
			}),
			want: "Os membros da familia Vidal Mendes são Andre, Manu, Matilda.",
		},
		{
			description: "Checa se o nome da familia for vazio",
			got: Say(Familia{
				Sobrenome: "",
				Membros:   []string{"Andre", "Manu", "Matilda"},
			}),
			want: "Os membros são Andre, Manu, Matilda.",
		},
		{
			description: "Sem nome da familia e sem membros",
			got: Say(Familia{
				Sobrenome: "",
				Membros:   []string{},
			}),
			want: "Não existe uma familia.",
		},
	}

	for _, st := range subtestes {
		t.Run(st.description, func(t *testing.T) {
			if st.got != st.want {
				t.Errorf("want: %v got: %v", st.want, st.got)
			}
		})
	}
}
