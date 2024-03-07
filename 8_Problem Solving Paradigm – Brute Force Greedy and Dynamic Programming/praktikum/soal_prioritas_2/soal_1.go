package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	cost := make([]int, n)
	cost[0] = 0

	for i := 1; i < n; i++ {
		cost[i] = math.MaxInt32
		for j := i - 1; j >= 0 && j >= i-2; j-- {
			jumpCost := int(math.Abs(float64(jumps[i] - jumps[j])))
			cost[i] = min(cost[i], cost[j]+jumpCost)
		}
	}

	return cost[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30

	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40

}
