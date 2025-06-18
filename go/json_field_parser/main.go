package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// type LogEntry struct {
// 	Time time.Time `json:"timestamp"`
// 	Message string `json:"message"`
// }

// func parseLogs(logFileName string) {
// 	file, err := os.Open(logFilename)
// 	if err != nil {
// 		fmt.Printf("Failed to open log file %s: %s\n", logFilename, err)
// 	}

// 	defer file.Close()

// 	var logEntry LogEntry
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		err := json.Unmarshal([]byte(scanner.Text()), &logEntry)
// 		if err != nil {
// 			fmt.Println("Failed to read line:", err)
// 			continue
// 		}

// 		fmt.Println(logEntry)
// 	}
// }

type LogLine struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Level     string    `json:"level"`
}

type Log []byte

type nginxLog struct {
	time.Time
}

func (nl *nginxLog) UnmarshalJSON(b []byte) {

}

func parseLog(log *[]byte) (*LogLine, error) {
	var logLine LogLine

	fmt.Println(string(*log))
	valid := json.Valid(*log)
	if !valid {
		return nil, fmt.Errorf("Not valid json")
	}

	json.Unmarshal(*log, &logLine)

	updateLevel(&logLine)

	return &logLine, nil
}

func updateLevel(ll *LogLine) {
	if ll.Level == "" {
		// * logic for parsing message and updating
		msg := strings.Fields(ll.Message)
		fmt.Println(ll.Message)

		var validHTTPCode = regexp.MustCompile(`^[1-5][0-9][0-9]$`)
		var validInfo = regexp.MustCompile(`^[2-3]`)
		var validWarn = regexp.MustCompile(`^[4]`)
		var validError = regexp.MustCompile(`^[5]`)
		if validHTTPCode.MatchString(msg[3]) {
			if validInfo.MatchString(msg[3]) {
				ll.Level = "Info"
				return
			}
			if validWarn.MatchString(msg[3]) {
				ll.Level = "Warn"
				return
			}
			if validError.MatchString(msg[3]) {
				ll.Level = "Error"
				return
			}
		}
	}
}

func main() {
	// 172.17.0.1 - - [06/Aug/2024:14:55:37 +0000] "GET / HTTP/1.1" 200 615 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"
	s := "GET /api/v1/endpoint HTTP/1.1"
	ss := strings.Fields(s)
	m := ss[0]

	fmt.Println(m)

	logByte := []byte(`{"timestamp": "06/Aug/2024:14:55:37 +0000", "message": "GET / HTTP/1.1 200 615 - Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"}`)
	out, err := parseLog(&logByte)
	if err != nil {
		fmt.Println("Error parsing log:", err)
	}
	fmt.Println(out)
}
