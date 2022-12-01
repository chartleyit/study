package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

func sumArray(a [3]int) int {
	r := 0
	for _, n := range a {
		r += n
	}
	return r
}

func updateCalories(c int, m [3]int) ([3]int) {
	for i := 0; i < len(m); i++ {
		if c >  m[i] {
			m[i], c = c, m[i]
		}
	}
	return m
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var maxCalories [3] int
	var calories int
	for fileScanner.Scan() {
		c, _ := strconv.Atoi(fileScanner.Text())
		calories += c
		if fileScanner.Text() == "" {
			if calories > maxCalories[2] {
				// func for managing list
				maxCalories = updateCalories(calories, maxCalories)
			}
			calories = 0
		}
	}

	readFile.Close()

	fmt.Println(maxCalories)
	fmt.Println(sumArray(maxCalories))
}
