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

func main() {
	// input is a single line
	input := readFile("input.part1")
	// input := readFile("input.sample")
	// input := readFile("input.sample2")
	for _, x := range input {
		fmt.Println(x)
	}
}
