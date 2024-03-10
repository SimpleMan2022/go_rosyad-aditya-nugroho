package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go PrimesAndPower(100, &wg)
	wg.Wait()
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimesAndPower(n int, wg *sync.WaitGroup) {
	chPrimes := make(chan int)
	chPowers := make(chan int)
	defer wg.Done()

	go func() {
		defer close(chPrimes)
		for i := 1; i <= n; i++ {
			if isPrime(i) {
				chPrimes <- i
			}
		}
	}()
	go func() {
		defer close(chPowers)
		for i := 1; i <= n; i++ {
			if isPrime(i) {
				chPowers <- i * i
			}
		}
	}()

	fmt.Print("Bilangan Prima 1 - 100: ")
	for prime := range chPrimes {
		fmt.Print(prime, " ")
	}

	fmt.Print("\n")
	fmt.Print("Bilangan Kuadrat Prima 1 - 100: ")
	for power := range chPowers {
		fmt.Print(power, " ")
	}

	fmt.Print("\n")
	fmt.Print("Total Bilangan kuadrat Prima 1 - 100: ")
	var total int
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			total += i * i
		}
	}
	fmt.Print(total)
}
