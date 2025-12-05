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

	zeroDialCount := 0
	currentDial := INITIAL_DIAL

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

		remainingDistance := distance % TOTAL_DIALS

		var calculatedDial int
		switch direction {
		case LEFT_DIRECTION:
			calculatedDial = currentDial - remainingDistance
			if calculatedDial >= 0 {
				currentDial = calculatedDial
			} else {
				currentDial = TOTAL_DIALS + calculatedDial
			}
		case RIGHT_DIRECTION:
			calculatedDial = currentDial + remainingDistance
			if calculatedDial < TOTAL_DIALS {
				currentDial = calculatedDial
			} else {
				currentDial = remainingDistance - (TOTAL_DIALS - currentDial)
			}
		}

		if currentDial == 0 {
			zeroDialCount++
		}
	}

	password := zeroDialCount
	fmt.Println("password:", password)
}
