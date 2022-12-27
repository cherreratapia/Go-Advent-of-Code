package aoc8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkLeftSide(row []int, y int) int {

	lookingFor := row[y]
	visibleTrees := 0
	for i := y - 1; i >= 0; i-- {
		if row[i] >= lookingFor {
			visibleTrees++
			break
		}
		visibleTrees++
	}

	return visibleTrees
}

func checkRightSide(row []int, y int) int {
	lookingFor := row[y]
	visibleTrees := 0

	for i := y + 1; i < len(row); i++ {

		if row[i] >= lookingFor {
			visibleTrees++
			break
		}
		visibleTrees++
	}

	return visibleTrees
}

func checkTopSide(treeMap [][]int, x int, y int) int {
	lookingFor := treeMap[x][y]
	visibleTrees := 0
	for i := x - 1; i >= 0; i-- {

		if treeMap[i][y] >= lookingFor {
			visibleTrees++
			break
		}
		visibleTrees++
	}
	return visibleTrees
}

func checkBottomSide(treeMap [][]int, x int, y int) int {
	lookingFor := treeMap[x][y]
	visibleTrees := 0
	for i := x + 1; i < len(treeMap); i++ {
		if treeMap[i][y] >= lookingFor {
			visibleTrees++
			break
		}
		visibleTrees++
	}
	return visibleTrees
}

func checkRow(treeMap [][]int, x int, y int) int {
	row := treeMap[x]
	return checkLeftSide(row, y) * checkRightSide(row, y)
}

func checkColumn(treeMap [][]int, x int, y int) int {
	return checkTopSide(treeMap, x, y) * checkBottomSide(treeMap, x, y)

}

func TreetopTreeHouse() {

	treeMap := [][]int{}
	visibleTrees := 0
	f, err := os.Open("./TreetopTreeHouse/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		heightsStr := strings.Split(line, "")
		for _, heightsStr := range heightsStr {
			if value, error := strconv.Atoi(heightsStr); error == nil {
				row = append(row, value)
			}
		}
		treeMap = append(treeMap, row)
	}

	fmt.Printf("Tree map %v\n", treeMap)
	for i := 1; i < len(treeMap)-1; i++ {
		for j := 1; j < len(treeMap[i])-1; j++ {
			scenicScore := checkRow(treeMap, i, j) * checkColumn(treeMap, i, j)
			fmt.Printf("x: %d, y: %d height: %d Scenic score: %d\n", i, j, treeMap[i][j], scenicScore)
			if scenicScore > visibleTrees {
				visibleTrees = scenicScore
			}
		}
	}
	fmt.Printf("Visible trees: %d\n", visibleTrees)
}
