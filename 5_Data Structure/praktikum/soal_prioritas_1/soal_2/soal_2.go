package main

import (
	"fmt"
)

func main() {
	fmt.Println(primeSum([]int{1, 2, 4, 5, 12, 19, 30})) // 26
	fmt.Println(primeSum([]int{45, 17, 44, 67, 11}))     // 95
	fmt.Println(primeSum([]int{15, 12, 9, 10}))          // 0
}

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}

	if number == 2 || number == 3 || number == 5 || number == 7 {
		return true
	}

	if number%2 == 0 || number%3 == 0 || number%5 == 0 || number%7 == 0 {
		return false
	}

	return true
}

func primeSum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		if isPrime(number) {
			sum += number
		}
	}
	return sum
}
