package main

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

func main() {
	kata := "kasur"
	printReverse(kata)
}

func reverse(kata string) string {
	splittedKata := strings.Split(kata, "")
	slices.Reverse(splittedKata)
	kataBaru := strings.Join(splittedKata, "")
	return kataBaru
}

func printReverse(kata string) {
	ch := make(chan string)
	go func() {
		defer close(ch)
		kataBaru := reverse(kata)
		for i := 1; i <= len(kataBaru); i++ {
			ch <- kataBaru[:i]
			time.Sleep(3 * time.Second)
		}
	}()

	for char := range ch {
		fmt.Println(char)
	}

}
