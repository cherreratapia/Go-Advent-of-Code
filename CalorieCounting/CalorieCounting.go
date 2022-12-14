package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func CalorieCounting() {
	f, err := os.Open("./CalorieCounting/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	caloriesCount := []int{}
	var caloriesSum int = 0

	for scanner.Scan() {
		line := scanner.Text()
		calories, _ := strconv.Atoi(line)
		if line == "" {
			caloriesCount = append(caloriesCount, caloriesSum)
			caloriesSum = 0
		}
		caloriesSum += calories
	}
	if !sort.IntsAreSorted(caloriesCount) {
		sort.Slice(caloriesCount, func(i, j int) bool {
			return caloriesCount[i] > caloriesCount[j]
		})
	}
	mostCalories := caloriesCount[0] + caloriesCount[1] + caloriesCount[2]
	fmt.Println("Most three calories: %v", mostCalories)
}
