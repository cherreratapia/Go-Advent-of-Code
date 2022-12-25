package aoc7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const MOVE_FOLDER string = "MOVE_FOLDER"
const LIST_FOLDER string = "LIST_FOLDER"

type fileType struct {
	name string
	size int64
}

func isCommand(line string) bool {
	return strings.Contains(line, "$")
}

func getCommand(line string) string {
	if strings.Contains(line, "cd") {
		return MOVE_FOLDER
	}
	if strings.Contains(line, "ls") {
		return LIST_FOLDER
	}
	return ""
}

func getFolderName(line string) string {
	splitted := strings.Split(line, " cd ")
	return splitted[1]
}

func getFile(line string) fileType {
	splitted := strings.Split(line, " ")
	size, _ := strconv.ParseInt(splitted[0], 10, 64)
	return fileType{splitted[1], size}
}

func NoSpaceLeftOnDevice() {

	listing := false
	stack := []string{}
	const TOTAL_SPACE_DISK = 70000000
	const UNUSED_SPACE = 30000000

	// map [folderPath] int
	sizePerFolder := map[string]int64{}
	f, err := os.Open("./NoSpaceLeftOnDevice/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if isCommand(line) {
			cmd := getCommand(line)
			listing = cmd == LIST_FOLDER

			if cmd == MOVE_FOLDER {
				folderName := getFolderName(line)
				if folderName == ".." {
					stack = stack[:len(stack)-1]
				} else {
					stackLen := len(stack)
					if stackLen == 0 {
						stack = append(stack, folderName)
						if _, ok := sizePerFolder[folderName]; !ok {
							sizePerFolder[folderName] = 0
						}
					} else {
						newPath := stack[stackLen-1] + folderName
						stack = append(stack, newPath)
						if _, ok := sizePerFolder[newPath]; !ok {
							sizePerFolder[newPath] = 0
						}
					}
				}
			}
			fmt.Printf("Stack ATM: %v\n", stack)
		} else {
			if listing {
				file := getFile(line)
				for _, path := range stack {
					if value, exists := sizePerFolder[path]; exists {
						sizePerFolder[path] = value + file.size
					}
				}
			}
		}

	}

	fmt.Printf("Size per folder %v\n", sizePerFolder)

	sizes := []int64{}
	for _, value := range sizePerFolder {
		sizes = append(sizes, value)
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] < sizes[j]
	})
	fmt.Printf("Sizes  %v\n", sizes)
	SPACE_AVAILABLE := TOTAL_SPACE_DISK - sizes[len(sizes)-1]
	SPACE_NEEDED := UNUSED_SPACE - SPACE_AVAILABLE
	for _, value := range sizes {
		if value >= SPACE_NEEDED {
			fmt.Printf("Space needed: %d\nValue: %d\n", SPACE_NEEDED, value)
			break
		}
	}

}
