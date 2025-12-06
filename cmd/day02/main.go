package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INPUT_FILE_PATH = "input.txt"

func main() {
	file, err := os.Open(INPUT_FILE_PATH)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer file.Close()

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
	}

	var invalidIds []string
	var invalidTotal int
	for r := range strings.SplitSeq(line, ",") {
		ids := strings.Split(r, "-")
		firstId, err := strconv.Atoi(ids[0])
		if err != nil {
			fmt.Printf("failed to convert string to integer: %v", err)
		}
		lastId, err := strconv.Atoi(ids[len(ids)-1])
		if err != nil {
			fmt.Printf("failed to convert string to integer: %v", err)
		}

		for i := firstId; i <= lastId; i++ {
			id := strconv.Itoa(i)
			if len(id)%2 == 0 {
				seqLength := len(id) / 2
				if id[:seqLength] == id[seqLength:] {
					invalidIds = append(invalidIds, id)
					invalidTotal += i
				}
			}
		}
	}

	fmt.Println("invalidIds:", invalidIds)
	fmt.Println("invalidTotal:", invalidTotal)
}
