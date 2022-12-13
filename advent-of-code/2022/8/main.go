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

// left and top need to be reversed
func scenicScore(p int, side []int) (int, []int) {
	c := 0
	v := []int{}

	if side == nil {
		return c, v
	}

	// fmt.Printf("\tview: %+v\n\t", side)
	for _, x := range side {
		if p > x {
			// fmt.Printf("%d ", x)
			v = append(v, x)
			c++
		} else {
			// fmt.Printf("%d EOL", x)
			v = append(v, x)
			c++
			break
		}
	}
	// fmt.Printf("\n")

	return c, v
}

func myReverse(l []int) []int {
	local := make([]int, len(l))
	copy(local, l)
	return local
}

func isVisible(v int, side1 []int, side2 []int) bool {
	// for part 2 we want to do similar things but not sort them
	// and want to iterate over the list
	sortedside1 := make([]int, len(side1))
	copy(sortedside1, side1)
	sort.SliceStable(sortedside1, func(i, j int) bool {
		return sortedside1[i] > sortedside1[j]
	})
	if v > sortedside1[0] {
		// fmt.Printf("side1: tree %v, is taller than %v, from %+v\n", v, sortedside1, sortedside1)
		return true
	}

	sortedside2 := make([]int, len(side2))
	copy(sortedside2, side2)
	sort.SliceStable(sortedside2, func(i, j int) bool {
		return sortedside2[i] > sortedside2[j]
	})
	if v > sortedside2[0] {
		// fmt.Printf("side2: tree %v, is taller than %v, from %+v\n", v, sortedside2, sortedside2)
		return true
	}

	return false
}

func main() {
	// input is a single line
	input := readFile("input.part1")
	// input := readFile("input.sample")
	// input := readFile("input.sample2")
	rows := make(map[int][]int, len(input))
	columns := make(map[int][]int, len(input[0]))
	maxScore := 0

	for r, x := range input {
		fmt.Println(x)
		strTrees := strings.Split(x, "")
		for c, strTree := range strTrees {
			tree, _ := strconv.Atoi(strTree)
			rows[r] = append(rows[r], tree)
			columns[c] = append(columns[c], tree)
		}
	}

	// fmt.Printf("rows: %+v\n", rows)
	// fmt.Printf("columns: %+v\n", columns)

	visible := (len(rows) * 2) + (len(columns) * 2) - 4
	fmt.Printf("Base visible: %d\n", visible)
	for x := 0; x < len(rows); x++ {
		for y := 0; y < len(columns); y++ {
			tree := columns[y][x]
			left, right := rows[x][:y], rows[x][y+1:]
			up, down := columns[y][:x], columns[y][x+1:]
			// fmt.Printf("x: %d, y: %d\n", x, y)
			// fmt.Printf("\trow: %v, %v\n", rows[x], tree)
			// fmt.Printf("\t\tleft: %+v, right: %+v\n", left, right)
			// fmt.Printf("\tcolumn: %v, %v\n", columns[y], tree)
			// fmt.Printf("\t\tup: %+v, down: %+v\n", up, down)

			if (x > 1 && y > 1) &&
				(x < len(rows)-1 && y < len(columns)-1) {
				if isVisible(tree, left, right) {
					// fmt.Printf("\t[%d,%d]: %v is visible in row\n", x, y, tree)
					visible++
				} else if isVisible(tree, up, down) {
					// fmt.Printf("\t[%d,%d]: %v is visible in column\n", x, y, tree)
					visible++
				}
			}

			// reverse
			u := myReverse(up)
			upScore, upV := scenicScore(tree, u)
			downScore, downV := scenicScore(tree, down)
			// reverse
			lt := myReverse(left)
			leftScore, leftV := scenicScore(tree, lt)
			rightScore, rightV := scenicScore(tree, right)
			score := upScore * downScore * leftScore * rightScore
			if score > maxScore {
				fmt.Printf("[%d,%d] %v\n", x, y, tree)
				// fmt.Printf("\trow: %v, %v\n", rows[x], tree)
				// fmt.Printf("\t\tleft: %+v, right: %+v\n", left, right)
				// fmt.Printf("\tcolumn: %v, %v\n", columns[y], tree)
				// fmt.Printf("\t\tup: %+v, down: %+v\n", up, down)
				fmt.Printf("\tup:\t%+v %d\n", u, len(u))
				fmt.Printf("\tupV:\t%+v %d\n", upV, len(upV))
				fmt.Printf("\tdown:\t%+v %d\n", down, len(down))
				fmt.Printf("\tdownV:\t%+v %d\n", downV, len(downV))
				fmt.Printf("\tleft:\t%+v %d\n", lt, len(lt))
				fmt.Printf("\tleftV:\t%+v %d\n", leftV, len(leftV))
				fmt.Printf("\tright:\t%+v %d\n", right, len(right))
				fmt.Printf("\trightV:\t%+v %d\n", rightV, len(rightV))
				fmt.Printf("\tup: %d, down: %d, left: %d, right: %d\n", upScore, downScore, leftScore, rightScore)
				fmt.Printf("\tscore: %d\n", score)
				fmt.Printf("\t%d is greater than %d\n", score, maxScore)
				maxScore = score
			}
		}
	}

	fmt.Printf("Field size: %d x %d\n", len(rows), len(columns))

	fmt.Printf("Visible trees: %v\n", visible)

	fmt.Printf("Max scenic score: %v\n", maxScore)
}
