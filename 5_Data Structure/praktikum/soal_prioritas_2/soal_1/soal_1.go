package main

import "fmt"

func main() {
	fmt.Println(spinSlice([]int{1, 2, 3, 4, 5}))    // [1,5,2,4,3]
	fmt.Println(spinSlice([]int{6, 7, 8}))          // [6,8,7]
	fmt.Println(spinSlice([]int{8, 5, 4, 3, 2, 1})) // [8,1,5,2,4,3]
}

func spinSlice(numbers []int) []int {
	n := len(numbers)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = numbers[i/2]
		} else {
			result[i] = numbers[n-(i/2)-1]
		}
	}
	return result
}
