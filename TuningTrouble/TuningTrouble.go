package aoc6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isMarker(word map[string]int) bool {
	for _, count := range word {
		if count > 1 {
			return false
		}
	}
	return true
}

func removeCountLetter(letter string, word map[string]int) map[string]int {
	value, exists := word[letter]
	if exists && value > 1 {
		word[letter] = value - 1
	}
	if exists && value == 1 {
		delete(word, letter)
	}
	return word
}

func TuningTrouble() {
	f, err := os.Open("./TuningTrouble/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	position := 0
	for scanner.Scan() {
		countLetters := map[string]int{}
		word := ""
		line := scanner.Text()
		for i, element := range line {
			letter := string(element)

			if i > 13 {
				if isMarker(countLetters) {
					position = i
				} else {
					countLetters = removeCountLetter(string(word[0]), countLetters)
					word = word[1:14]
				}
			}
			if position != 0 {
				break
			}
			word += letter
			value, exists := countLetters[letter]
			if exists {
				countLetters[letter] = value + 1
			} else {
				countLetters[letter] = 1
			}
		}
	}
	fmt.Printf("Marker: %v\n", position)
}
