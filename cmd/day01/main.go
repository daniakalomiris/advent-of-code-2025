package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const INPUT_FILE_PATH = "input.txt"
const TOTAL_DIALS = 100
const INITIAL_DIAL = 50
const LEFT_DIRECTION = "L"
const RIGHT_DIRECTION = "R"

func main() {
	file, err := os.Open(INPUT_FILE_PATH)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	resultPart1 := part1(lines)
	fmt.Println("password part 1:", resultPart1)

	resultPart2 := part2(lines)
	fmt.Println("password part 2:", resultPart2)
}

func part1(lines []string) int {
	pointCount := 0

	dial := INITIAL_DIAL

	for _, l := range lines {
		direction := l[:1]
		distance, err := strconv.Atoi(l[1:])
		if err != nil {
			fmt.Printf("failed to convert string to integer: %v", err)
		}

		remainingDistance := distance % TOTAL_DIALS

		var calculatedDial int
		switch direction {
		case LEFT_DIRECTION:
			calculatedDial = dial - remainingDistance
			if calculatedDial >= 0 {
				dial = calculatedDial
			} else {
				dial = TOTAL_DIALS + calculatedDial
			}
		case RIGHT_DIRECTION:
			calculatedDial = dial + remainingDistance
			if calculatedDial < TOTAL_DIALS {
				dial = calculatedDial
			} else {
				dial = remainingDistance - (TOTAL_DIALS - dial)
			}
		}

		if dial == 0 {
			pointCount++
		}
	}

	return pointCount
}

func part2(lines []string) int {
	pointCount := 0
	clickCount := 0

	dial := INITIAL_DIAL

	for _, l := range lines {
		direction := l[:1]
		distance, err := strconv.Atoi(l[1:])
		if err != nil {
			fmt.Printf("failed to convert string to integer: %v", err)
		}

		fullRotationCount := distance / 100
		if fullRotationCount >= 1 {
			clickCount += fullRotationCount
		}

		remainingDistance := distance % TOTAL_DIALS

		var calculatedDial int
		switch direction {
		case LEFT_DIRECTION:
			calculatedDial = dial - remainingDistance
			if calculatedDial >= 0 {
				dial = calculatedDial
			} else {
				if dial != 0 {
					clickCount++
				}
				dial = TOTAL_DIALS + calculatedDial
			}
		case RIGHT_DIRECTION:
			calculatedDial = dial + remainingDistance
			if calculatedDial < TOTAL_DIALS {
				dial = calculatedDial
			} else {
				dial = remainingDistance - (TOTAL_DIALS - dial)
				if dial != 0 {
					clickCount++
				}
			}
		}

		if dial == 0 {
			pointCount++
		}
	}

	return clickCount + pointCount
}
