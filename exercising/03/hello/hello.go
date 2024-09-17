package hello

import (
	"fmt"
	"strings"
)

type Familia struct {
	Sobrenome string
	Membros   []string
}

func Say(f Familia) string {
	s := strings.Join(f.Membros, ", ")

	if f.Sobrenome == "" {
		if len(f.Membros) < 1 {
			return fmt.Sprintln("Não existe uma familia.")
		} else {
			return fmt.Sprintf("Os membros são %v.", s)
		}
	}

	return fmt.Sprintf("Os membros da familia %v são %v.", f.Sobrenome, s)
}
