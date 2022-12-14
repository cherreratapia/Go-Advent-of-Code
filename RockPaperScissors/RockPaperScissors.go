package aoc2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func _GetResultPoints(result string) int {
	resultPoints := map[string]int{
		"LOSE": 0,
		"DRAW": 3,
		"WIN":  6,
	}
	return resultPoints[result]
}

func _GetShapePoints(shape string) int {
	shapePoints := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	return shapePoints[shape]
}

func _InstructionToDo(opponent string, cryptedInstruction string) int {
	UncryptInstruction := map[string]string{
		"X": "LOSE",
		"Y": "DRAW",
		"Z": "WIN",
	}

	Posibilities := map[string]map[string]string{
		"A": {"WIN": "B", "DRAW": "A", "LOSE": "C"},
		"B": {"WIN": "C", "DRAW": "B", "LOSE": "A"},
		"C": {"WIN": "A", "DRAW": "C", "LOSE": "B"},
	}

	instruction := UncryptInstruction[cryptedInstruction]
	playerShape := Posibilities[opponent][instruction]
	resultPoints := _GetResultPoints(instruction)
	shapePoints := _GetShapePoints(playerShape)
	return resultPoints + shapePoints

}

func _SplitInstructions(line string) (string, string) {
	slice := strings.Split(line, " ")
	return slice[0], slice[1]
}

// func _PlayerPoints(opponent string, player string) int {
// 	points := map[string]int{
// 		"X": 1,
// 		"Y": 2,
// 		"Z": 3,
// 	}
// 	/*
// 		A = X = ROCK
// 		B = Y = PAPER
// 		C = Z = SCISSORS
// 	*/
// 	WonPoints := 6
// 	DrawPoints := 3

// 	if opponent == "A" && player == "Y" || opponent == "B" && player == "Z" || opponent == "C" && player == "X" {
// 		return points[player] + WonPoints
// 	}
// 	if opponent == "A" && player == "X" || opponent == "B" && player == "Y" || opponent == "C" && player == "Z" {
// 		return points[player] + DrawPoints
// 	}
// 	return points[player]
// }

func RockPaperScissors() {

	f, err := os.Open("./RockPaperScissors/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var points int = 0
	for scanner.Scan() {
		instructions := scanner.Text()
		opponentInstruction, expectedResult := _SplitInstructions(instructions)
		points += _InstructionToDo(opponentInstruction, expectedResult)
	}
	fmt.Println("Total points: %v", points)

}
