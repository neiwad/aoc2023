package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var total = 0
var numberBuffer = ""
var isStringAlphabetic = regexp.MustCompile(`[a-z,A-Z,0-9,.]`).MatchString

func convertStringToTotal(s []string, numberBuffer string) int {
	specialCharacters := regexp.MustCompile(`[a-z,A-Z,0-9,.]`).ReplaceAllString(strings.Join(s, ""), "")
	number, err := strconv.Atoi(numberBuffer)
	if err == nil {
		return len(specialCharacters) * number
	}
	return 0
}

func checkNeighbours(i int, j int, k int, lines [][]string) {
	// Check top
	if i > 0 {
		left := j
		right := k
		if left > 0 {
			left = j - 1
		}
		if right < len(lines[i]) {
			right = k + 1
		}
		topString := lines[i-1][left:right]
		// Check number of special characters in topString
		total += convertStringToTotal(topString, numberBuffer)
	}

	// Check bottom
	if i < len(lines)-1 {
		left := j
		right := k
		if left > 0 {
			left = j - 1
		}
		if right < len(lines[i]) {
			right = k + 1
		}
		bottomString := lines[i+1][left:right]
		// Check number of special characters in topString
		total += convertStringToTotal(bottomString, numberBuffer)
	}

	// Check left
	if j > 0 {
		left := j - 1
		leftString := strings.Split(lines[i][left], "")
		// Check number of special characters in topString
		total += convertStringToTotal(leftString, numberBuffer)
	}

	// Check right
	if k < len(lines[i])-1 {
		right := k
		rightString := strings.Split(lines[i][right], "")
		// Check number of special characters in topString
		total += convertStringToTotal(rightString, numberBuffer)
	}
}

func main() {
	readFile, err := os.Open("../../input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var lines [][]string

	for fileScanner.Scan() {
		lineString := fileScanner.Text()
		lines = append(lines, strings.Split(lineString, ""))
	}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {

			// If it's a number
			_, err := strconv.Atoi(lines[i][j])
			if err == nil {
				for k := j; k < len(lines[i]); k++ {
					number, err := strconv.Atoi(lines[i][k])
					// If it's a number
					if err == nil {
						numberBuffer = fmt.Sprintf("%s%s", numberBuffer, strconv.Itoa((number)))

						// Check if it's the end of the line
						if k == len(lines[i])-1 {
							checkNeighbours(i, j, k, lines)
							numberBuffer = ""
							j = k
							break
						}

					} else {
						// End of number sequence
						// Check all neighbours to find special characters

						checkNeighbours(i, j, k, lines)
						numberBuffer = ""
						j = k
						break

					}
				}

			}
		}
	}

	fmt.Println("Day 3, Part 1: ", total)
}
