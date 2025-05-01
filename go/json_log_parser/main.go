package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Log struct {
	timestamp time.Time `json:"timestamp"` //type for a time stamp
	level     string
	msg       string `json:"message"`
}

func main() {
	fmt.Println("hello from main")

	// open file
	// parse/unmarshal json
	//

	logMsg := Log{}
	json.NewDecoder([]bytes).Decode(&logMsg)
	// disallow unknown fields

	json.Unmarshal([]bytes, logMsg)
}
