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
var baseCards []string
var cards []string

func calculatePower(winnings []string, numbers []string) int {
	count := 0
	for i := 0; i < len(numbers); i++ {
		if slices.Contains(winnings, numbers[i]) {
			count++
		}
	}

	return count

}

func splitCard(card string) ([]string, []string) {
	parts := strings.Split(card, ":")
	cardNumbers := strings.Split(parts[1], "|")
	winnings := strings.Fields(cardNumbers[0])
	numbers := strings.Fields(cardNumbers[1])

	return winnings, numbers
}

func processCards(cards []string) {
	for i := 0; i < len(cards); i++ {
		winnings, numbers := splitCard(cards[i])
		power := calculatePower(winnings, numbers)
		total += int(math.Pow(2, float64(power-1)))
	}
}

func addCardsCopy(i int, card string, baseCards []string) {
	winnings, numbers := splitCard(card)
	power := calculatePower(winnings, numbers)

	cards = append(cards, card)

	if power > 0 {
		for j := i + 1; j < len(baseCards) && j < i+1+power; j++ {
			tempCard := baseCards[j]
			addCardsCopy(j, tempCard, baseCards)
		}
	}
}

func main() {
	readFile, err := os.Open("../../input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lineString := fileScanner.Text()
		baseCards = append(baseCards, lineString)
	}

	for i := 0; i < len(baseCards); i++ {
		addCardsCopy(i, baseCards[i], baseCards)
	}

	fmt.Println("Total Cards: ", len(cards))

	fmt.Println("Day 4, Part 2: ", total)
}
