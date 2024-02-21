package main

import (
	"fmt"
	"math"
)

func main() {
	const phi float64 = 3.14
	var (
		r float64 = 14
		t float64 = 20
	)
	fmt.Println("Program Menghitung Volume Tabung")

	// rumus luas volume tabung = phi * r^2 * t
	luas := phi * math.Pow(r, 2) * t
	fmt.Println(luas)

}
