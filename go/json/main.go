package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

type Request struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Greeting string `json:"greeting"`
}

func marshalJSON() {
	p := Person{Name: "Chris", Age: 42}
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(string(jsonData))
}

func unmarshalJSON() {
	jsonData := []byte(`{"Name":"Chris","Age":42}`)
	isValid := json.Valid(jsonData)
	if !isValid {
		fmt.Printf("Error: invalid json %s\n", jsonData)
		return
	}

	var p Person
	err := json.Unmarshal(jsonData, &p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(p)
}

func encodeJSON() {
	p := Person{Name: "Chris", Age: 42}

	f, err := os.Create("my.json")
	if err != nil {
		fmt.Println("Error: encoding data", err)
	}

	encoder := json.NewEncoder(f)
	err = encoder.Encode(p)
	if err != nil {
		fmt.Println("Error: encoding data", err)
	}
}

func decodeJSON() {
	jsonData := `{"Name":"Chris","Age":42}`
	reader := strings.NewReader(jsonData)
	decoder := json.NewDecoder(reader)

	var p Person
	err := decoder.Decode(&p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(p)
}

func decodeJSONapi() {
	fmt.Println("decoding json api response")
	reqBody := Request{Name: "Bob", Age: 36}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Post("http://example.com/api/greet", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var respBody Response
	json.NewDecoder(resp.Body).Decode(&respBody)
	fmt.Println(respBody.Greeting)
}

func unmarshalJSONfile() {
	filename := "input.json"
	fmt.Println("Unmarshaling JSON file:", filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file:", err)
	}

	defer file.Close()

	var p Person
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// this should be a complete line of json for a json log
		err := json.Unmarshal([]byte(line), &p)
		if err != nil {
			log.Println("Failed to parse line:", err)
			continue
		}
		fmt.Println(p)
	}
}

func main() {
	marshalJSON()
	unmarshalJSON()
	encodeJSON()
	decodeJSON()

	decodeJSONapi()
	unmarshalJSONfile()
}
