package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
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
	//
	// part2
	// overlap at all.
	// begin1 <= begin2 AND begin2 <= end1
	// begin1 <= end2 AND end2 <= end1
	input := readFile("input.part1")
	numOverlapInt := 0
	completeOverlapInt := 0

	for _, line := range input {
		ranges := strings.Split(line, ",")
		range1, range2 := strings.Split(ranges[0], "-"), strings.Split(ranges[1], "-")
		begin1, _ := strconv.Atoi(range1[0])
		end1, _ := strconv.Atoi(range1[1])
		begin2, _ := strconv.Atoi(range2[0])
		end2, _ := strconv.Atoi(range2[1])

		if begin1 <= begin2 && end1 >= end2 {
			numOverlapInt += 1
			completeOverlapInt += 1
			fmt.Printf("Range1: %v, contains Range2: %v\n", range1, range2)
		} else if begin2 <= begin1 && end2 >= end1 {
			numOverlapInt += 1
			completeOverlapInt += 1
			fmt.Printf("Range2: %v, contains Range1: %v\n", range2, range1)
		} else if begin1 <= begin2 && begin2 <= end1 {
			numOverlapInt += 1
			fmt.Printf("Range2: %v, starts within Range1: %v\n", range2, range1)
		} else if begin1 <= end2 && end2 <= end1 {
			numOverlapInt += 1
			fmt.Printf("Range2: %v, starts within Range1: %v\n", range2, range1)
		}
	}
	fmt.Println("Part1: Number of complete overlaps: ", completeOverlapInt)
	fmt.Println("Part2: Number of all overlaps int: ", numOverlapInt)
}