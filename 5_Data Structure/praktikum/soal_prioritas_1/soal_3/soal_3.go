package main

import (
	"fmt"
	"slices"
)

func main() {
	data1 := []float64{1, 1, 2, 5, 6, 8, 12, 4, 5, 5, 5, 8, 9}

	fmt.Println("sum: ", sum(data1))       // 71.00
	fmt.Println("mean: ", mean(data1))     // 5.46
	fmt.Println("median: ", median(data1)) // 5.00
	fmt.Println("mode: ", mode(data1))     // 5.00

	data2 := []float64{6, 7, 1, 11, 8, 12, 6, 12, 1, 7, 8, 10, 14}

	fmt.Println("sum: ", sum(data2))       // 103.00
	fmt.Println("mean: ", mean(data2))     // 7.92
	fmt.Println("median: ", median(data2)) // 8.00
	fmt.Println("mode: ", mode(data2))     // 1.00
}

func sum(data []float64) string {
	var result float64
	for _, char := range data {
		result += char
	}
	return fmt.Sprintf("%.2f", result)

}

func mean(data []float64) string {
	var result float64
	for _, char := range data {
		result += char
	}
	lenData := float64(len(data))
	return fmt.Sprintf("%.2f", result/lenData)
}

func median(data []float64) string {
	lenData := len(data)
	slices.Sort(data)
	if lenData%2 == 0 {
		medianVal := (lenData + 1) / 2
		return fmt.Sprintf("%.2f", data[medianVal-1])
	} else {
		x1 := lenData / 2
		x2 := (lenData / 2) + 1
		medianVal := int((data[x1-1] + data[x2-1]) / 2)
		return fmt.Sprintf("%.2f", data[medianVal])
	}
}

func mode(data []float64) string {
	freq := map[float64]int{}
	var maxFreq int
	var modes []float64

	for _, char := range data {
		freq[char]++
		if freq[char] > maxFreq {
			maxFreq = freq[char]
			modes = []float64{char}
		} else if freq[char] == maxFreq {
			modes = append(modes, char)
		}
	}
	return fmt.Sprintf("%.2f", modes[0])

}
