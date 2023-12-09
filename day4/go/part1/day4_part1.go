package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

var total = 0

func main() {
	readFile, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		count := 0

		lineString := fileScanner.Text()
		parts := strings.Split(lineString, ":")
		cardNumbers := strings.Split(parts[1], "|")
		winnings := strings.Fields(cardNumbers[0])
		numbers := strings.Fields(cardNumbers[1])

		//fmt.Println(winnings, numbers)

		for i := 0; i < len(numbers); i++ {
			if slices.Contains(winnings, numbers[i]) {
				count++
			}
		}

		total += int(math.Pow(2, float64(count-1)))

	}

	fmt.Println("Day 4, Part 1: ", total)
}
