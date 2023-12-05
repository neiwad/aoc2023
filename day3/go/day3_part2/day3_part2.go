package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Gear struct {
	x      int
	y      int
	number int
}

var total = 0
var numberBuffer = ""
var isStringAlphabetic = regexp.MustCompile(`[a-z,A-Z,0-9,.]`).MatchString
var gears = []Gear{}

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
		topString := strings.Join(lines[i-1][left:right], "")

		// Get index of gear
		gearIndex := strings.Index(topString, "*")

		if gearIndex != -1 {
			number, _ := strconv.Atoi(numberBuffer)
			gears = append(gears, Gear{x: gearIndex + left, y: i - 1, number: number})
		}
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
		bottomString := strings.Join(lines[i+1][left:right], "")

		// Get index of gear
		gearIndex := strings.Index(bottomString, "*")

		if gearIndex != -1 {
			number, _ := strconv.Atoi(numberBuffer)
			gears = append(gears, Gear{x: gearIndex + left, y: i + 1, number: number})
		}
	}

	// Check left
	if j > 0 {
		left := j - 1
		leftString := strings.Join(strings.Split(lines[i][left], ""), "")

		gearIndex := strings.Index(leftString, "*")
		if gearIndex != -1 {
			number, _ := strconv.Atoi(numberBuffer)
			gears = append(gears, Gear{x: left, y: i, number: number})
		}
	}

	// Check right
	if k < len(lines[i])-1 {
		right := k
		rightString := strings.Split(lines[i][right], "")

		gearIndex := strings.Index(strings.Join(rightString, ""), "*")
		if gearIndex != -1 {
			number, _ := strconv.Atoi(numberBuffer)
			gears = append(gears, Gear{x: right, y: i, number: number})
		}
	}
}

type Key struct {
	x int
	y int
}

type GearTotal struct {
	gear    Gear
	numbers []int
}

func multiplyGears(gears []Gear) map[Key]GearTotal {
	gearCount := make(map[Key]GearTotal)

	// Count occurrences of each unique combination of x and y
	for _, gear := range gears {
		key := Key{gear.x, gear.y}

		gt, exists := gearCount[key]
		if exists {
			// If the key exists, update the numbers slice
			gt.numbers = append(gt.numbers, gear.number)
		} else {
			// If the key doesn't exist, create a new entry in the map
			gt = GearTotal{gear, []int{gear.number}}
		}

		// Assign the updated or new value back to the map
		gearCount[key] = gt
	}

	return gearCount
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

	res := multiplyGears(gears)

	fmt.Println(res)
	for _, gear := range res {
		if len(gear.numbers) == 2 {
			total += gear.numbers[0] * gear.numbers[1]
		}
	}

	fmt.Println("Day 3, Part 2: ", total)
}
