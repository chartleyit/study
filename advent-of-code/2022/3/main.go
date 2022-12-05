package main

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
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

func priority(item rune) int {
	// "unicode/utf8"
	// lower ord a - (value)
	offset, begin := 1, 'a'
	// upper ord a - value + 27
	if unicode.IsUpper(item) {
		offset, begin = 27, 'A'
	} 

	diff := item - begin
	p := offset + int(diff)
	fmt.Printf("%v - %v = %v + %v = %v [%c]\n", begin, item, diff, offset, p, item)

	return p
}

func getDuplicate(sack string) rune {
	l := len(sack) / 2
	r := []rune(sack)
	comp1 := r[l:]
	comp2 := r[:l]
	
	items := make(map[rune]bool)
	for _, i := range comp1 {
		items[i] = true
	}

	for _, x := range comp2 {
		if items[x] {
			return x
		}
	}
	return 'a'
}

func main() {
	input := readFile("input.part2")
	total := 0
	for i, x := range input {
		// every 3 packs is a different group
		group := i / 3
		fmt.Printf("Group %v: ", group)
		missPacked := getDuplicate(x)
		total += priority(missPacked)
	}

	fmt.Println(total)
}