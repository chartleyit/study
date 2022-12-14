
 package main

 import (
	"fmt"
	"bufio"
	"os"
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

 func result(opp int, me int) int {
	var r int

	// case statement (honestly this code be static for simplicity sake since the cases are limited)
	// A X tie 3 rock v rock 0
	// A Y win 6 rock v paper 1
	// B X lost 0 paper v rock -1
	// A Z loss 0 rock v scissors -2
	// C X win 6 scissors vs rock 2
	switch diff := opp - me; diff {
		case 1:
			r = 0
		case 0:
			r = 3
		case -1:
			r = 6
		case -2:
			r = 0
		case 2:
			r = 6
	}

	return r
 }

 func maxThree(i int) int {
	if i > 2 {
		return 0
	} else if i < 0 {
		return 2
	}
	return i
 }

 func decidePick(oppPick string, result string) string {
	// choices := []string{"A", "B", "C"}
	choices := "ABC"
	myChoices := []string{"X", "Y", "Z"}
	i := strings.Index(choices, oppPick)
	switch result {
	case "X": // lose
		i = maxThree(i - 1)

	case "Z": // win
		i = maxThree(i + 1)
	default: // draw
		// nothing
	}

	return myChoices[i]
 }

 func main() {
	total := 0

	Score := map[string]int{
		"X": 1, // rock // lose
		"Y": 2, // paper // draw
		"Z": 3, // scissor // win
	}

	Opp := map[string]int{
		"A": 1, // rock
		"B": 2, // paper
		"C": 3, // scissor
	}

	lines := readFile("input.txt")
	for _, line := range lines {
		round := strings.Split(line, " ")
		oppChoice, expectedResult := round[0], round[1]

		// change things up here
		// remap XYZ into what "would" get that result
		// provide that data to the next step for the same calculations
		myChoice := decidePick(oppChoice, expectedResult)

		result := result(Opp[oppChoice], Score[myChoice])
		roundScore := Score[myChoice] + result
		total += roundScore

		// fmt.Printf("%v %v + %v = %v %v\n", round, Score[myChoice], result, roundScore, total)
	}


	fmt.Println("Total: ", total)
 }