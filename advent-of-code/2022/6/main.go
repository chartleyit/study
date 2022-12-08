package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(file string) []string {
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			fileLines = append(fileLines, line)
		}
	}

	readFile.Close()

	return fileLines
}

func signalParser(r []rune, l int) {
	msgLength := len(r)
	pos := 0
	signalFound := false
	// TODO need to think about this logic this might end up short
	for signalFound != true {
		if pos >= msgLength-l {
			break
		}
		signalMarker := r[pos:(pos + l)]
		fmt.Println(pos, string(signalMarker))
		if isUnique(signalMarker) {
			signalFound = true
			fmt.Printf("Found signal marker: %s, at position %d\n", string(signalMarker), pos+l)
		}
		pos++
		// fmt.Printf("pos, end %d, %d\n", pos, msgLength-l)
	}
}

func isUnique(r []rune) bool {
	counts := make(map[rune]int)
	for _, x := range r {
		fmt.Printf("%+v", string(x))
		counts[x]++
		if counts[x] > 1 {
			return false
		}
	}
	fmt.Printf("%+v\n", counts)
	return true
}

func main() {
	// input is a single line
	input := readFile("input.part1")
	markerSize := 4

	// how many characters need to be parsed before the first start
	// of packet is received
	// start of packet is 4 non-repeating characters
	signalParser([]rune(input[0]), markerSize)
	fmt.Println("message length: ", len([]rune(input[0])))
}
