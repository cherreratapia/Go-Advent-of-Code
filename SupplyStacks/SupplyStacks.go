package aoc5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getColumns(raw []string) map[int][]string {
	inputLen := len(raw)
	columns := strings.ReplaceAll(raw[inputLen-1], " ", "")
	result := map[int][]string{}
	for index := range columns {
		result[index+1] = []string{}
	}
	return result
}

func fillCrates(rawCrates []string, columns map[int][]string) map[int][]string {

	for row := len(rawCrates) - 1; row >= 0; row-- {
		for index, unicode := range rawCrates[row] {
			letter := string(rune(unicode))
			if letter != " " && letter != "[" && letter != "]" {
				stack := ((index) / 4) + 1
				columns[stack] = append(columns[stack], letter)
			}
		}
	}
	return columns
}

func formatInstructions(raw []string) []int {
	result := []int{}
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	for i := range raw {
		elements := re.FindAllString(raw[i], -1)
		for _, element := range elements {
			value, _ := strconv.Atoi(element)
			result = append(result, value)
		}
	}
	return result
}

func execInstructions(instructions []int, columns map[int][]string) map[int][]string {
	for i := 0; i < len(instructions); i += 3 {
		move := instructions[i]
		src := instructions[i+1]
		to := instructions[i+2]
		cratesToMove := []string{}
		for j := 0; j < move; j++ {
			cratesToMove = append([]string{columns[src][len(columns[src])-1]}, cratesToMove...)
			columns[src] = columns[src][:len(columns[src])-1]
		}
		columns[to] = append(columns[to], cratesToMove...)
	}
	return columns
}

func getTopCrates(columns map[int][]string) string {
	result := ""
	lenColumns := len(columns)
	for i := 0; i < lenColumns; i++ {
		result += columns[i+1][len(columns[i+1])-1]
	}
	return result
}

func SupplyStacks() {
	f, err := os.Open("./SupplyStacks/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	rawCrates := []string{}
	columns := map[int][]string{}
	rawInstructions := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && !strings.Contains(line, "move") {
			rawCrates = append(rawCrates, line)
			result, _ := regexp.MatchString("[1-9]", line)
			if result {
				columns = getColumns(rawCrates)
				rawCrates = rawCrates[:len(rawCrates)-1]
				columns = fillCrates(rawCrates, columns)
			}
		}
		if line != "" && strings.Contains(line, "move") {
			rawInstructions = append(rawInstructions, line)
		}
	}
	instructions := formatInstructions(rawInstructions)
	columns = execInstructions(instructions, columns)
	topCrates := getTopCrates(columns)
	fmt.Printf("Top crates are: %v\n", topCrates)
}
