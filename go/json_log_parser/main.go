package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Log struct {
	Timestamp time.Time `json:"timestamp"` //type for a time stamp
	Level     string    `json:"level"`
	Msg       string    `json:"msg"`
}

func main() {
	logFilename := "log.json"

	file, err := os.Open(logFilename)
	if err != nil {
		fmt.Printf("Failed to open log file %s: %s\n", logFilename, err)
	}

	defer file.Close()

	var logEntry Log
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err := json.Unmarshal([]byte(scanner.Text()), &logEntry)
		if err != nil {
			fmt.Println("Failed to read line:", err)
			continue
		}

		fmt.Println(logEntry)
	}
}
