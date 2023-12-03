package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numbers = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func day1_part2() {
	readFile, err := os.Open("../input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	total := 0
	for fileScanner.Scan() {

		line := fileScanner.Text()
		newLine := ""

		// Loop through each character in the string
		for i := 0; i < len(line); i++ {

			char := line[i]

			// Check if the character is a number
			if _, err := strconv.Atoi(string(char)); err == nil {
				newLine += string(char)
			}

			// Loop through each number
			for y, number := range numbers {
				// Check if the number is in the string
				if i+len(number) < len(line)+1 {
					stringToCheck := line[i : i+len(number)]
					if string(stringToCheck) == number {
						newLine += strconv.Itoa(y + 1)
					}
				}

			}

		}

		if len(newLine) != 0 {
			// Construct the number
			number := fmt.Sprintf("%s%s", string(newLine[0]), string(newLine[len(newLine)-1]))

			// Convert to int and add to total
			res, err := strconv.Atoi(number)
			if err == nil {
				total += res
			}
		}

	}

	readFile.Close()
	fmt.Println("Part2: ", total) // Need to be 54728
}
