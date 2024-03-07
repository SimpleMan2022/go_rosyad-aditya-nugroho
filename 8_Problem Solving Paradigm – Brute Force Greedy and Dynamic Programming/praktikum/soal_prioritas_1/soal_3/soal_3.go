package main

import "fmt"

var memo = make(map[int]int)

func fibonacci(number int) int {
	if number <= 1 {
		return number
	}

	if memo[number] != 0 {
		return memo[number]
	}

	memo[number] = fibonacci(number-1) + fibonacci(number-2)
	return memo[number]

}

func main() {

	fmt.Println(fibonacci(0)) // 0

	fmt.Println(fibonacci(2)) // 1

	fmt.Println(fibonacci(9)) // 34

	fmt.Println(fibonacci(10)) // 55

	fmt.Println(fibonacci(12)) // 144

}
