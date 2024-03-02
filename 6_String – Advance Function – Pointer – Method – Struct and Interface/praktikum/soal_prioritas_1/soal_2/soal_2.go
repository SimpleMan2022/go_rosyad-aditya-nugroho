package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(spinString("alta"))    // aalt
	fmt.Println(spinString("alterra")) // aalrtre
	fmt.Println(spinString("sepulsa")) // saesplu
}

func spinString(sentence string) string {
	var result []string
	splitSentence := strings.Split(sentence, "")
	n := len(splitSentence)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result = append(result, splitSentence[i/2])
		} else {
			result = append(result, splitSentence[n-(i/2)-1])
		}
	}
	return strings.Join(result, "")
}
