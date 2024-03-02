package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(groupPalindrome([]string{"katak", "civic", "kasur", "rusak"}))                  // [[katak, civic], kasur, rusak]
	fmt.Println(groupPalindrome([]string{"racecar", "seru", "kasur", "civic", "bilik", "kak"})) // [[racecar, civic, kak], seru, kasur, bilik]
	fmt.Println(groupPalindrome([]string{"masuk", "civic", "hahah", "garam"}))                  // [[civic, hahah], masuk, garam]
}

func isPalindrom(word string) bool {
	splittedWord := strings.Split(word, "")
	slices.Reverse[[]string](splittedWord)
	if strings.Join(splittedWord, "") != word {
		return false
	}
	return true
}

func groupPalindrome(words []string) []any {
	var result []any
	var palindrom []any
	for _, word := range words {
		if isPalindrom(word) {
			palindrom = append(palindrom, word)
		} else {
			result = append(result, word)
		}
	}

	if len(palindrom) > 0 {
		result = append([]any{palindrom}, result...)
		return result
	}
	return result
}
