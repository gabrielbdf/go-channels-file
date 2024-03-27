package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	getWithChannels()

	elapsed := time.Since(start)
	log.Printf("getJson took %s", elapsed)

}

func getWithChannels() {
	myChannel := make(chan map[string]any, 100)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			myChannel <- getJson("https://dog-api.kinduff.com/api/facts")
		}()
	}
	go func() {
		wg.Wait()
		close(myChannel)
	}()

	myResponses := []map[string]any{}
	for value := range myChannel {
		myResponses = append(myResponses, value)
		fmt.Println(value)
	}

	file, err := os.Create("responses.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&myResponses)
	if err != nil {
		log.Fatal(err)
	}
}
func getWithoutChannels() {
	myResponses := []map[string]any{}
	for i := 0; i < 10; i++ {
		myResponses = append(myResponses, getJson("https://dog-api.kinduff.com/api/facts"))
	}
	for _, value := range myResponses {
		fmt.Println(value)
	}
	file, err := os.Create("responses.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&myResponses)
	if err != nil {
		log.Fatal(err)
	}

}

func getJson(link string) map[string]any {

	var response map[string]any

	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
	return response

}
