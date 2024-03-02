package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pickVocals("alterra academy"))     // aea aae
	fmt.Println(pickVocals("sepulsa mantap jiwa")) // eua aa ia
	fmt.Println(pickVocals("kopi susu"))           // oi uu
}

func pickVocals(sentence string) string {
	var temp []string
	splitted := strings.Split(sentence, " ")
	for _, word := range splitted {
		var tempWord []string
		charSplit := strings.Split(word, "")
		for _, char := range charSplit {
			if char == "a" || char == "e" || char == "i" || char == "o" || char == "u" {
				tempWord = append(tempWord, char)
			}
		}
		temp = append(temp, strings.Join(tempWord, ""))
	}
	return strings.Join(temp, " ")
}
