package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	mychannel := make(chan bool, 5)
	sem := make(chan bool, 2)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		sem <- true
		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()
			fmt.Printf("Inicio processamento: %d\n", i)
			time.Sleep(3 * time.Second)
			fmt.Printf("Final processamento: %d\n", i)
			mychannel <- true
		}(i)
	}
	for i := 0; i < 5; i++ {
		<-mychannel
	}

	wg.Wait()
}
