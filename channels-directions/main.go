package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func ping(pings chan<- string, msg string, wait int) {
	time.Sleep(time.Duration(wait) * time.Second)
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	file, err := os.OpenFile("./filename.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	start := time.Now()
	pings := make(chan string)
	pongs := make(chan string)

	go ping(pings, "passed message 1", 1)
	go ping(pings, "passed message 4", 4)
	go ping(pings, "passed message 3", 3)
	go ping(pings, "passed message 2", 2)
	go pong(pings, pongs)
	go pong(pings, pongs)
	go pong(pings, pongs)
	go pong(pings, pongs)

	for i := 0; i < 4; i++ {
		if _, err := file.WriteString("teste\n"); err != nil {
			log.Fatal(err)
		}
		fmt.Println(<-pongs)
	}

	elapsed := time.Since(start)
	log.Printf("getJson took %s", elapsed)

}
