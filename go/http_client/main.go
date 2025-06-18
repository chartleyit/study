package main

/*
	inspired by this youtube video
	https://jsonplaceholder.typicode.com/todos/1
	https://www.youtube.com/watch?v=e4UkT-EyHQo
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todos struct {
	UID       int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	r, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal("Error: failed to get url", err)
	}

	if r.StatusCode != http.StatusOK {
		log.Fatal("Something went wrong.  HTTP Response", r.StatusCode)
	}

	var todos Todos
	err = json.NewDecoder(r.Body).Decode(&todos)
	if err != nil {
		log.Fatal("failed to decode json response body", err)
	}

	fmt.Println(todos)

	// rb, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	log.Fatal("Failed to read body", err)
	// }
	// fmt.Println(string(rb))
}
