package main

import (
	"fmt"
	"os"
	"bufio"
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
	// In how many assignment pairs does one range fully contain the other?
	// inputs include 2 ranges
	// need to parse the beginning and end of the range and identify when
	// begin1 < begin2 and end1 > end2
	// or
	// begin2 < begin1 and end2 > end1
	input := readFile("input.part1")
	numOverlap := 0

	for _, line := range input {
		ranges := strings.Split(line, ",")
		range1, range2 := strings.Split(ranges[0], "-"), strings.Split(ranges[1], "-")
		if range1[0] <= range2[0] && range1[1] >= range2[1] {
			numOverlap += 1
			fmt.Printf("Range1: %v, contains Range2: %v\n", range1, range2)
		} else if range2[0] <= range1[0] && range2[1] >= range1[1] {
			numOverlap += 1
			fmt.Printf("Range2: %v, contains Range1: %v\n", range2, range1)
		}
	}
	fmt.Println("Number of overlaps: ", numOverlap)
}