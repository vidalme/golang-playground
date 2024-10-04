package main

import (
	"fmt"
	"sync"
	"time"
)

func minutim(t int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Me da %d segundos\n", t)
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Printf("Obrigado por aguardar %d segundos.\n", t)
}

func main() {
	var wg sync.WaitGroup
	count := 3
	for i := 0; i < count; i++ {
		wg.Add(1)
		go minutim(i*2, &wg)
	}

	wg.Wait()
	fmt.Println("interacao final, finaliza ao mesmo tempo que a ultima contagem do waiting group")

}
