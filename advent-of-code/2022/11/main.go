package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Monkey struct {
	startingItems []int
	operation	string
	test		Test
	inspections	int
}

// !! issue with setting the ifTrue and ifFalse
// div gets set correctly but not the otherse
type Test struct {
	div		int
	ifTrue	int
	ifFalse int
}

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

func parseMonkey(m *Monkey, instructions string) {
	i := strings.Split(instructions, ":")
	// fmt.Printf("%+v\n", i)
	// TODO replace this with a case statement
	// starting
	if strings.Contains(i[0], "Starting") {
		fmt.Println("\tChecking for starting items... ")
		i[1] = strings.ReplaceAll(i[1], " ", "")
		itemsStr := strings.Split(i[1], ",")
		items := make([]int, len(itemsStr))
		for n, str := range itemsStr {
			items[n], _ = strconv.Atoi(str)
		}
		m.startingItems = items
	}
	// operation
	if strings.Contains(i[0], "Operation") {
		fmt.Println("\tChecking for operations...")
		m.operation = i[1]
	}
	// test
	if strings.Contains(i[0], "Test") {
		// This assumes the test is divisble vs some other calculation
		fmt.Println("\tChecking for test...")
		opStr := strings.Split(i[1], " ")
		op, _ := strconv.Atoi(opStr[3])
		m.test.div = op
	}
	// if true/if false
	if strings.Contains(i[0], "If false") {
		fmt.Println("\tChecking for ifFalse...")
		fStr := strings.Split(i[1], " ")
		m.test.ifFalse, _ = strconv.Atoi(fStr[4])
	}
	// if true/if false
	if strings.Contains(i[0], "If true") {
		fmt.Println("\tChecking for ifTrue...")
		tStr := strings.Split(i[1], " ")
		m.test.ifTrue, _ = strconv.Atoi(tStr[4])
	}
}

func newMonkey() *Monkey {
	return &Monkey{}
}

// throw
func throw(current int, target int, monkeys []*Monkey) {
	// delete
	item, currentItems := monkeys[current].startingItems[0], monkeys[current].startingItems[1:]
	monkeys[current].startingItems = currentItems

	// append
	monkeys[target].startingItems = append(monkeys[target].startingItems, item)
}

// test
func test(value int, div int, ifTrue int, ifFalse int) int {
	if value % div == 0 {
		return ifTrue
	} else {
		return ifFalse
	}
}

// operation
func operation(operation string, item int) int {
	op := strings.Split(operation, " ")
	var first int
	var second int
	var err error
	if op[3] == "old" {
		first = item
	} else {
		first, err = strconv.Atoi(op[3])
		if err != nil {
			fmt.Println(err)
		}
	}
	if op[5] == "old" {
		second = item
	} else {
		second, err = strconv.Atoi(op[5])
		if err != nil {
			fmt.Println(err)
		}
	}

	if op[4] == "+" {
		return first + second
	} else if op[4] == "*" {
		return first * second
	} else if op[4] == "/" {
		return first / second
	} else if op[4] == "-" {
		return first - second
	}

	return item
}

// inspect
func inspect(item int, m *Monkey, monkeys []*Monkey, current int) {
		// fmt.Printf("%d ", item)
		// operation
		v := operation(m.operation, item)
		// fmt.Println("before", item, "after", v)
		// relief / 3 round down
		v = v / 3
		// fmt.Println("relief", v)
		// test
		target := test(v, m.test.div, m.test.ifTrue, m.test.ifFalse)
		// throw
		throw(current, target, monkeys)
		m.inspections += 1
}

func bubbleSort(arr []int) []int {
	for i:=0; i< len(arr) -1; i++ {
		for j:=0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	// input is a single line
	input := readFile("input.part1")
	// input := readFile("input.sample")
	// input := readFile("input.sample2")
	monkeys := []*Monkey{}
	var monkeyPtr *Monkey
	for _, x := range input {
		// fmt.Println(x)
		// initialize monkey if it's a monkey
		if strings.Contains(x, "Monkey") {
			monkeyPtr = newMonkey()
			monkeys = append(monkeys, monkeyPtr)
			fmt.Println("New", x, monkeyPtr)
		} else {
			parseMonkey(monkeyPtr, x)
		}

		// else set attributes of monkey
	}


	for _, m := range monkeys {
		fmt.Println(*m)
	}
	// this creates a copy of monkeys
	// TODO do this 20 times
	for n := 0; n < 20; n++ {
		fmt.Println("Round", n)
		for i, m := range monkeys {
			fmt.Println("starting", i, m.startingItems)
			for _, item := range m.startingItems {
				inspect(item, m, monkeys, i)
			}
		}
		fmt.Printf("\n")
		for _, m := range monkeys {
			fmt.Printf("%d ", m.inspections)
		}
		fmt.Printf("\n")
	}

	monkeyBuisness := []int{}
	fmt.Println(monkeyBuisness)
	for _, m := range monkeys {
		fmt.Println(*m)
		monkeyBuisness = append(monkeyBuisness, m.inspections)
		fmt.Println(monkeyBuisness)
	}

	fmt.Println(monkeyBuisness)
	monkeyBuisness = bubbleSort(monkeyBuisness)
	fmt.Println(monkeyBuisness)
	fmt.Println("top 2 monkeys", monkeyBuisness[0], monkeyBuisness[1])
	fmt.Println("score", monkeyBuisness[0]*monkeyBuisness[1])
	// score 68900 is too low :-()
}
