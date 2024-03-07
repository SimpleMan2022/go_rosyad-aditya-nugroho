package main

import "fmt"

func convertRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var result string
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			result += romans[i]
			num -= vals[i]
		}
	}
	return result
}

func main() {
	fmt.Println(convertRoman(4))
	fmt.Println(convertRoman(9))
	fmt.Println(convertRoman(23))
	fmt.Println(convertRoman(2021))
	fmt.Println(convertRoman(1646))
}
