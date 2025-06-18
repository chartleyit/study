package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// * create log structure
type LogEntry struct {
	Timestamp StdTime `json:"timestamp"`
	Level     string  `json:"level"`
	Event     string  `json:"event"`
}

// * create custom log structure
type StdTime struct {
	time.Time
}

// * create type handler
func (st *StdTime) UnmarshalJSON(b []byte) error {
	var err error
	str := strings.Trim(string(b), `"`)
	layouts := []string{
		time.RFC3339,
		"2006-01-02 15:04:05",
		"02-Jan-2006 15:04:05",
		time.ANSIC,
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, str); err == nil {
			st.Time = t
			return nil
		}
	}

	return fmt.Errorf("unknown timestamp format: %s", err)
}

func readLogs(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open file %s, %s\n", fileName, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry LogEntry
		err := json.Unmarshal([]byte(scanner.Text()), &entry)
		if err != nil {
			fmt.Println("Failed to read line:", err)
			continue
		}
		fmt.Println(entry)
	}
}

func main() {
	readLogs("log.json")
}
