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

func getDuplicate(comp1 []rune, comp2 []rune) (rune, []rune) {
	// to handle 3 lists of multiples this should return a list of runes
	// to no break previous logic just extending it with additional return
	var rList []rune
	var r rune
	
	items := make(map[rune]bool)
	for _, i := range comp1 {
		items[i] = true
	}

	for _, x := range comp2 {
		if items[x] {
			rList = append(rList, x)
			r = x
		}
	}
	return r, rList
}

func splitString(s string) ([]rune, []rune) {
	l := len(s) / 2
	r := []rune(s)
	comp1 := r[l:]
	comp2 := r[:l]

	return comp1, comp2
}

func findKey(rList1 []rune, rList2 []rune, rList3 []rune) rune {
	// iterate over 3 lists for intersection
	// run duplicates twice

	_, overLap := getDuplicate(rList1, rList2)
	key, _ := getDuplicate(overLap, rList3)
	
	return key
}

func main() {
	input := readFile("input.part2")
	total := 0
	keyTotal := 0
	errors := map[int]rune{}
	groups := map[int][]rune{}
	keys := map[int]rune{}
	for i, x := range input {
		// every 3 packs is a different group
		// find the value that is in all 3 strings
		group := i / 3
		fmt.Printf("Group %v: ", group)

		if i % 3 == 0 {
			keys[group] = findKey(groups[i-2], groups[i-1], groups[i])
			keyTotal += priority(keys[group])
		}
		
		c1, c2 := splitString(x)
		missPacked, _ := getDuplicate(c1, c2)
		errors[group] = missPacked
		total += priority(missPacked)
	}

	fmt.Printf("%+v\n", groups)
	fmt.Println("Total Errors: ", total)
	fmt.Println("Key Total: ", keyTotal)
}