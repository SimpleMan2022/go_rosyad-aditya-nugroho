package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getItem([]int{10, 10, 5, 30, 40, 10, 5}, 80)) // [40 30 10]
	fmt.Println(getItem([]int{50, 20, 10, 25, 25}, 100))      // [50 25 25]
}

func getItem(items []int, target int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i] > items[j]
	})

	result := []int{}
	for _, item := range items {
		if item <= target {
			result = append(result, item)
			target -= item
		}
		if target == 0 {
			break
		}
	}
	return result
}
