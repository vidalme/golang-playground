package main

import "fmt"

type reader interface {
	read() (int, error)
}

type writer interface {
	write() (int, error)
}

type readWriter interface {
	reader
	writer
}

type printer struct {
	name string
}

func (p printer) read() (int, error) {
	fmt.Println("------------")
	fmt.Println("metodo read")
	fmt.Println("------------")
	return 1, nil
}

func (p printer) write() (int, error) {
	fmt.Println("------------")
	fmt.Println("metodo write")
	fmt.Println("------------")
	return 1, nil
}

func main1() {
	p := printer{"hp"}
	var rw readWriter = p
	rw.read()
	rw.write()
}
