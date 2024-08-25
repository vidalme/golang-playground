package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("Ola mundo: %v\n Argumentos: %v\n", args, args[1:])
}
