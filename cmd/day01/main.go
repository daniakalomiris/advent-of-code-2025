package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const INPUT_FILE_PATH = "input.txt"
	const TOTAL_DIALS = 100
	const INITIAL_DIAL = 50
	const LEFT_DIRECTION = "L"
	const RIGHT_DIRECTION = "R"

	zeroPointCount := 0
	zeroClickCount := 0

	dial := INITIAL_DIAL

	file, err := os.Open(INPUT_FILE_PATH)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[:1]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("failed to convert string to integer: %v", err)
		}

		fullRotationCount := distance / 100
		if fullRotationCount >= 1 {
			zeroClickCount += fullRotationCount
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
					zeroClickCount++
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
					zeroClickCount++
				}
			}
		}

		if dial == 0 {
			zeroPointCount++
		}
	}

	password := zeroPointCount + zeroClickCount
	fmt.Println("password:", password)
}
