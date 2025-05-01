package main

import "fmt"

func twoSum(nums []int, target int) []int {
	prevMap := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		if _, containsKey := prevMap[diff]; containsKey {
			return []int{prevMap[diff], i}
		}
		prevMap[n] = i
	}
	fmt.Println(target)
	fmt.Println(nums)
	return []int{}
}

func main() {
	input := []int{2, 7, 11, 15}

	keyPair := twoSum(input, 9)

	fmt.Println(keyPair)
}
