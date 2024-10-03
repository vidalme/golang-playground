package main

import (
	"time"
)

func main() {

	go conectaBD()
	go pegaDados()
	go updateDados()

}

func conectaBD() {
	time.Sleep(1000 * time.Millisecond)
}

func pegaDados() {
	time.Sleep(800 * time.Millisecond)
}

func updateDados() {
	time.Sleep(2000 * time.Millisecond)
}
