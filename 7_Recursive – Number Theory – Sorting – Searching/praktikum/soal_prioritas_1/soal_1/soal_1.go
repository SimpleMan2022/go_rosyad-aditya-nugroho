package main

import "fmt"

func main() {
	fmt.Println(fibX(5))  // 12
	fmt.Println(fibX(10)) // 143
}

func fibX(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += fibonacci(i)
	}
	return sum
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
