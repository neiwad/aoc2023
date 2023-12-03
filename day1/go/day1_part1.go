package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func day1_part1() {
	readFile, err := os.Open("../input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	total := 0
	for fileScanner.Scan() {

		// Remove all non-numeric characters
		reg := regexp.MustCompile("[a-z, A-Z]")
		stringWithoutLetters := reg.ReplaceAllString(fileScanner.Text(), "")

		if len(stringWithoutLetters) != 0 {
			// Construct the number
			number := fmt.Sprintf("%s%s", string(stringWithoutLetters[0]), string(stringWithoutLetters[len(stringWithoutLetters)-1]))

			// Convert to int and add to total
			res, err := strconv.Atoi(number)
			if err == nil {
				total += res
			}
		}

	}

	readFile.Close()
	fmt.Println("Part1: ", total)
}
