package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func isVisible(v int, side1 []int, side2 []int) bool {
	// a tree is visible if there are no taller tress
	// between it and any edge
	// we can rule out 0 and len-1 for each row and column
	// since the outside edge itself is visible
	sortedside1 := make([]int, len(side1))
	copy(sortedside1, side1)
	sort.SliceStable(sortedside1, func(i, j int) bool {
		return sortedside1[i] > sortedside1[j]
	})
	if v > sortedside1[0] {
		fmt.Printf("side1: tree %v, is taller than %v, from %+v\n", v, sortedside1, sortedside1)
		return true
	}

	sortedside2 := make([]int, len(side2))
	copy(sortedside2, side2)
	sort.SliceStable(sortedside2, func(i, j int) bool {
		return sortedside2[i] > sortedside2[j]
	})
	if v > sortedside2[0] {
		fmt.Printf("side2: tree %v, is taller than %v, from %+v\n", v, sortedside2, sortedside2)
		return true
	}

	return false
}

func main() {
	// input is a single line
	input := readFile("input.part1")
	// input := readFile("input.sample")
	rows := make(map[int][]int, len(input))
	columns := make(map[int][]int, len(input[0]))

	for r, x := range input {
		strTrees := strings.Split(x, "")
		for c, strTree := range strTrees {
			tree, _ := strconv.Atoi(strTree)
			rows[r] = append(rows[r], tree)
			columns[c] = append(columns[c], tree)
		}
	}

	fmt.Printf("rows: %+v\n", rows)
	fmt.Printf("columns: %+v\n", columns)

	visible := (len(rows) * 2) + (len(columns) * 2) - 4
	fmt.Printf("Base visible: %d\n", visible)
	for x := 1; x < len(rows)-1; x++ {
		for y := 1; y < len(columns)-1; y++ {
			// fmt.Printf("rows: %+v\n", rows)
			// fmt.Printf("columns: %+v\n", columns)
			tree := columns[y][x]
			left, right := rows[x][:y], rows[x][y+1:]
			up, down := columns[y][:x], columns[y][x+1:]
			fmt.Printf("x: %d, y: %d\n", x, y)

			fmt.Printf("\trow: %v, %v\n", rows[x], tree)
			fmt.Printf("\t\tleft: %+v, right: %+v\n", left, right)
			fmt.Printf("\tcolumn: %v, %v\n", columns[y], tree)
			fmt.Printf("\t\tup: %+v, down: %+v\n", up, down)
			if isVisible(tree, left, right) {
				fmt.Printf("\t[%d,%d]: %v is visible in row\n", x, y, tree)
				visible++
			} else if isVisible(tree, up, down) {
				fmt.Printf("\t[%d,%d]: %v is visible in column\n", x, y, tree)
				visible++
			}
		}
	}

	fmt.Printf("rows: %+v\n", rows)
	fmt.Printf("columns: %+v\n", columns)
	fmt.Printf("Visible trees: %v", visible)
}
