package aoc4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	start  int
	finish int
}

func getPoints(elf string) Elf {
	elfAssignment := strings.Split(elf, "-")
	start, startErr := strconv.Atoi(elfAssignment[0])
	if startErr != nil {
		log.Fatal("Error parsing start: %v", startErr)
	}
	finish, finishErr := strconv.Atoi(elfAssignment[1])
	if finishErr != nil {
		log.Fatal("Error parsing finish: %v", startErr)
	}
	return Elf{start, finish}

}

func getAssignmentElf(instructions string) (Elf, Elf) {
	elves := strings.Split(instructions, ",")
	firstElf := getPoints(elves[0])
	secondElf := getPoints(elves[1])
	return firstElf, secondElf
}

func CampCleanup() {
	f, err := os.Open("./CampCleanup/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	rangeContained := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstElf, secondElf := getAssignmentElf(line)
		// if firstElf.start >= secondElf.start &&
		// 	firstElf.finish <= secondElf.finish || secondElf.start >= firstElf.start &&
		// 	secondElf.finish <= firstElf.finish {
		// 	rangeContained++
		// }
		if firstElf.start >= secondElf.start && firstElf.start <= secondElf.finish || firstElf.finish >= secondElf.start && firstElf.finish <= secondElf.finish ||
			secondElf.start >= firstElf.start && secondElf.start <= firstElf.finish || secondElf.finish >= firstElf.start && secondElf.finish <= firstElf.finish {
			rangeContained++
		}
	}
	fmt.Printf("Range contained %v\n", rangeContained)
}
