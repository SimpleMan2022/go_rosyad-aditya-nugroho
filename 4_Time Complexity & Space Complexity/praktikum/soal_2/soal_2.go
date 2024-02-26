package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		bagi := pow(x, n/2)
		return bagi * bagi
	}
	bagi := pow(x, (n-1)/2)
	return bagi * bagi * x
}

func main() {
	fmt.Println(pow(2, 5))
}
