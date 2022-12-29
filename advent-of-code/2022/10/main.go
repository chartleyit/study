package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	insSet := make([]int, len(input)*2)
	insIndex := 0

	reg := 1
	regTrack := 0

	for _, x := range input {

		insIndex += 1
		if strings.Contains(x, "addx") {
			fmt.Printf("%v: %v\n", insIndex, x)
			i, _ := strconv.Atoi(strings.Split(x, " ")[1])
			insIndex += 1
			insSet[insIndex] = i
		}

	}

	fmt.Println("================================")
	fmt.Printf("%+v\n", insSet)
	fmt.Println("================================")

	fmt.Println(regTrack)
	for i, x := range insSet {
		// fmt.Printf("%v: %v\n", i, x)
		if i == 0 {
			fmt.Println("skip first cycle")
		} else if (i+20)%40 == 0 {
			regTrack += reg * i
			fmt.Printf("%v: %v, %v\n", i, reg, regTrack)
		}
		reg += x
	}
	fmt.Printf("%v\n", regTrack)
}
