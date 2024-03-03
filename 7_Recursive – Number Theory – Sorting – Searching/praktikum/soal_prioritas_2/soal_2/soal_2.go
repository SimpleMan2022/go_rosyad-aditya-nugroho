package main

import "fmt"

func main() {
	primeFactors(20)   // 2, 2, 5
	primeFactors(75)   // 3, 5, 5
	primeFactors(1024) // 2, 2, 2, 2, 2, 2, 2, 2, 2, 2
}

func primeFactors(n int) {
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d", i)
			n /= i
			if n > 1 {
				fmt.Print(", ")
			} else {
				fmt.Println()
			}
			primeFactors(n)
			return
		}
	}
}
