package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Soal Eksplorasi : Anagram")
	kata := "kasur"
	kata2 := "rusak"

	if len(kata) != len(kata2) {
		fmt.Println("Bukan Anagram")
	}

	splitKata := strings.Split(kata, "")
	sort.Strings(splitKata)
	finalKata := strings.Join(splitKata, "")

	splitKata2 := strings.Split(kata2, "")
	finalKata2 := strings.Join(splitKata2, "")

	if finalKata != finalKata2 {
		fmt.Println("Bukan Anagram")
	} else {
		fmt.Println("Anagram")
	}

}
