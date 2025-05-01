// max subarray
package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []float64) float64 {
	var maxSub float64 = nums[0] // initialize the maxSub with the first value in nums
	var curSum float64 = 0       // set our current sum

	for _, n := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += n
		maxSub = math.Max(maxSub, curSum)
	}
	return maxSub
}

func main() {
	fmt.Println("Hello, world!")

	testNums := []float64{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	r := maxSubArray(testNums)
	fmt.Println(r)
}
