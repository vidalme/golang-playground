package interacao

func Repete(s string) string {
	ns := s
	for i := 0; i < 4; i++ {
		ns += s
	}
	return ns
}
