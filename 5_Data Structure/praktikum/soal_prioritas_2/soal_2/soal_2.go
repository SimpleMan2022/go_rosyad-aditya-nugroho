package main

import (
	"fmt"
)

func main() {

	fmt.Println(groupPrime([]int{2, 3, 4, 5, 6, 7, 8})) // [[2,3,5,7],4,6,8]
	fmt.Println(groupPrime([]int{15, 24, 3, 9, 5}))     // [[3,5],15,24,9]
	fmt.Println(groupPrime([]int{4, 8, 9, 12}))         // [4,8,9,12]
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

func groupPrime(numbers []int) []any {

	primes := []any{}
	res := []any{}
	for _, number := range numbers {
		if isPrime(number) {
			primes = append(primes, number)
		} else {
			res = append(res, number)
		}
	}

	if len(primes) > 0 {
		res = append([]any{primes}, res...)
		return res
	}
	return res
}
