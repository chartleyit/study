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

func praseStacks(vStack []string, n int) [][]rune {
	// count every x chars
	// 0123456789
	// [H] [M] [N] [Z] [M] [C] [M] [P] [P]
	//  x   x   x   x   x   x   x   x   x
	// 1 5 9 13 17 21 25 29 33
	// -1 % 4
	// 0 4 8 12 16
	// every 3 starting at 1
	// i + 3 % 3 = 0
	stacks := make([][]rune, n)

	for _, x := range vStack {
		runes := []rune(x)
		for i, r := range runes {
			if (i-1)%4 == 0 && r != ' ' {
				s := (i - 1) / 4
				// this needs to prepend in order to make it LIFO
				stacks[s] = append([]rune{r}, stacks[s]...)
			}
		}
	}

	return stacks
}

func parsInstructions(ins []string) [][]int {
	insSets := [][]int{}
	for _, i := range ins {
		set := strings.Split(i, " ")
		numMove, _ := strconv.Atoi(set[1])
		from, _ := strconv.Atoi(set[3])
		to, _ := strconv.Atoi(set[5])
		insSets = append(insSets, []int{numMove, from, to})
	}

	return insSets
}

func arrangeStacks(s [][]rune, i [][]int) [][]rune {
	for _, x := range i {
		// modify a slice
		// grab the last value on the stack and append it
		fmt.Printf("initial: num stacks: %d, stack sizes: ", len(s))
		for _, d := range s {
			fmt.Printf("%d ", len(d))
		}
		fmt.Printf("\n")
		fmt.Printf("instructions %v from %v to %v\n", x[0], x[1], x[2])
		from := x[1] - 1
		to := x[2] - 1
		move := s[from][len(s[from])-x[0]:]
		fmt.Printf("moving %d\n", len(move))
		if len(move) != x[0] {
			fmt.Printf("ERROR: incorrect number moved: %d should be %d\n", len(move), x[0])
		}
		// part 2 just stats multiples aren't reversed (1 by 1)
		// rev := reverseSlice(move)
		rev := move
		if len(rev) != x[0] {
			fmt.Printf("ERROR: incorrect number reversed moved: %d should be %d\n", len(rev), x[0])
		}
		fmt.Printf("before: %d %s %d %s\n", from+1, string(s[from]), to+1, string(s[to]))
		s[from] = s[from][:len(s[from])-x[0]]
		s[to] = append(s[to], rev...)
		fmt.Printf("after: %d %s %d %s\n", from+1, string(s[from]), to+1, string(s[to]))

		// this is dumb but what are the lengths of all of our lists
		fmt.Printf("post: num stacks: %d, stack sizes: ", len(s))
		for _, d := range s {
			fmt.Printf("%d ", len(d))
		}
		fmt.Printf("\n")
	}

	return s
}

func reverseSlice(r []rune) []rune {
	runeCopy := make([]rune, len(r))
	copy(runeCopy, r)
	for i, j := 0, len(runeCopy)-1; i < j; i, j = i+1, j-1 {
		runeCopy[i], runeCopy[j] = runeCopy[j], runeCopy[i]
	}
	return runeCopy
}

func main() {
	// input := readFile("input.sample")
	input := readFile("input.part1")
	instructions := []string{}
	maxStack := 8 // maybe it's actually 8
	numStacks := 9
	// maxStack := 3
	// numStacks := 3
	stacks := []string{}

	for n, x := range input {
		fmt.Println("ITEM ", n)
		if n < maxStack {
			stacks = append(stacks, x)
			fmt.Println(x)
		}
		if n == maxStack {
			fmt.Printf("%v stacks\n", x)
		}
		if n > maxStack {
			instructions = append(instructions, x)
		}
	}

	parsedStacks := praseStacks(stacks, numStacks)
	// fmt.Printf("prased stacks: %+v\n", parsedStacks)
	parsedIns := parsInstructions(instructions)
	// fmt.Printf("parsed ins: %+v\n", parsedIns)
	arrangedStacks := arrangeStacks(parsedStacks, parsedIns)
	// fmt.Printf("arrangedStacks %+v", arrangedStacks)

	fmt.Printf("Last value from each stack: ")
	for _, s := range arrangedStacks {
		fmt.Printf("%+v", string(s[len(s)-1]))
	}
	fmt.Printf("\n")
	for _, s := range arrangedStacks {
		fmt.Printf("%+v\n", string(s))
	}
}
