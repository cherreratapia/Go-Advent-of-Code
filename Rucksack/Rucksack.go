package aoc3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func getItemsByCompartments(items string) (string, string) {
	half := len(items) / 2
	firstCompartment := items[:half]
	secondCompartment := items[half:]
	return firstCompartment, secondCompartment
}

func getASCIICode(letter rune) int {
	char := int(letter)
	char -= 64
	return char
}

func getDecimalValue(letter rune) int {
	oppositeLetter := letter
	isLower := unicode.IsLower(letter)
	if isLower {
		oppositeLetter = unicode.ToUpper(oppositeLetter)
		return getASCIICode(oppositeLetter)
	} else {
		oppositeLetter = unicode.ToLower(oppositeLetter)
		return getASCIICode(oppositeLetter) - 6
	}
}

func Rucksack() {

	f, err := os.Open("./Rucksack/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	duplicateItems := 0
	letters := map[rune]int{}
	for i := 1; scanner.Scan(); i++ {
		rucksackItems := scanner.Text()

		if len(letters) == 0 {
			for _, letter := range rucksackItems {
				letters[(letter)] = 1
			}
		} else {
			for letter := range letters {
				if !strings.ContainsRune(rucksackItems, letter) {
					delete(letters, letter)
				}
			}
		}
		if i%3 == 0 && i > 0 {
			for letter := range letters {
				duplicateItems += getDecimalValue(letter)
			}
			letters = map[rune]int{}
		}

	}
	fmt.Printf("Final value: %v\n", duplicateItems)
}

func firstStar(rucksackItems string) {
	duplicatedLetters := map[string]int{}
	duplicateItems := 0
	firstCompartment, secondCompartment := getItemsByCompartments(rucksackItems)
	for _, letter := range secondCompartment {
		index := strings.Index(firstCompartment, string(letter))
		if index != -1 {
			_, exists := duplicatedLetters[string(letter)]
			if !exists {
				duplicatedLetters[string(letter)] = 1
				duplicateItems += getDecimalValue(letter)
			}

		}
	}
	duplicatedLetters = map[string]int{}
}
