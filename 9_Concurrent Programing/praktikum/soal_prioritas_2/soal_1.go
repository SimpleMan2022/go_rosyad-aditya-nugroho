package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	file, err := os.Open("document.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var wg sync.WaitGroup
	var mutex sync.Mutex
	scanner := bufio.NewScanner(file)
	wordMap := make(map[string]int)
	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Split(text, " ")

		wg.Add(1)
		go func(words []string) {
			defer wg.Done()
			CountFrequency(words, wordMap, &mutex)
		}(words)
	}

	wg.Wait()

	for word, freq := range wordMap {
		fmt.Printf("%s: %d\n", word, freq)
	}
}

func CountFrequency(word []string, wordMap map[string]int, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	for _, char := range word {
		wordMap[char]++
	}
}
