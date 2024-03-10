package main

import (
	"fmt"
)

func main() {
	powerCh := make(chan int)
	resultCh := make(chan string)
	defer close(powerCh)
	defer close(resultCh)

	for i := 1; i <= 100; i++ {
		if isComposite(i) {
			go PowerComposite(i, powerCh, resultCh)
			fmt.Printf("Composite: %d^2 (%d) = %s\n", i, <-powerCh, <-resultCh)
		}
	}
}

func isComposite(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return true
		}
	}
	return false
}

func PowerComposite(n int, powerCh chan int, resultCh chan string) {

	power := n * n
	powerCh <- power
	if power%2 == 0 {
		resultCh <- "Ping"
	} else {
		resultCh <- "Pong"
	}

}
