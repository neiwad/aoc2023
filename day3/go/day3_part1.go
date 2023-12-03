package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var isStringAlphabetic = regexp.MustCompile(`[a-z,A-Z,0-9,.]`).MatchString

func day3_part1() {
	readFile, err := os.Open("../input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var lines [][]string
	total := 0

	for fileScanner.Scan() {

		lineString := fileScanner.Text()
		line := strings.Split(lineString, "")

		// For each character in the line
		for i := 0; i < len(line); i++ {

			_, err = strconv.Atoi(line[i])
			// If it's a number
			if err == nil {
				number := ""
				// Keep going until we hit a non-number
				for j := i; j < len(line); j++ {
					_, err = strconv.Atoi(line[j])
					// If it's a number, add it to the number string
					if err == nil {
						number += string(line[j])
					} else {
						// If it's not a number, add the number to the line
						for k := i; k < j; k++ {
							line[k] = string(number)
						}
						i += len(number) - 1
						break
					}
				}

			}
		}
		lines = append(lines, line)
	}

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			// If it's a special character
			if !isStringAlphabetic(lines[i][j]) {
				fmt.Println("Special character: ", lines[i][j])

				// Top left
				if i > 0 && j > 0 {
					neightbour, err := strconv.Atoi(lines[i-1][j-1])
					if err == nil {
						fmt.Println(neightbour)
						total += neightbour
					}
				}
				// Top center
				if i > 0 {
					neightbour, err := strconv.Atoi(lines[i-1][j])
					if err == nil {
						fmt.Println(neightbour)
						total += neightbour
					}
				}
				// Top right
				if i > 0 && j < len(lines[i])-1 {
					neightbour, err := strconv.Atoi(lines[i-1][j+1])
					if err == nil {
						fmt.Println(neightbour)
						total += neightbour
					}
				}
				// Center left
				if j > 0 {
					neightbour, err := strconv.Atoi(lines[i][j-1])
					if err == nil {
						fmt.Println(neightbour)
						total += neightbour
					}
				}
				// Center right
				if j < len(lines[i])-1 {
					neightbour, err := strconv.Atoi(lines[i][j+1])
					if err == nil {
						fmt.Println(neightbour)
						total += neightbour
					}
				}
				// Bottom left
				if i < len(lines)-1 && j > 0 {
					neightbour, err := strconv.Atoi(lines[i+1][j-1])
					if err == nil {
						fmt.Println(neightbour)
						total += neightbour
					}
				}
				// Bottom center
				if i < len(lines)-1 {
					neightbour, err := strconv.Atoi(lines[i+1][j])
					if err == nil {
						fmt.Println(neightbour)
						total += neightbour
					}
				}
				// Bottom right
				if i < len(lines)-1 && j < len(lines[i])-1 {
					neightbour, err := strconv.Atoi(lines[i+1][j+1])
					if err == nil {
						fmt.Println(neightbour)
						total += neightbour
					}
				}

			}
		}
	}

	fmt.Println("Day 3, Part 1: ", total)
}
