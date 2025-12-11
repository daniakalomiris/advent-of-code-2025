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

	resultPart1 := part1(line)
	fmt.Println("invalid total part 1:", resultPart1)

	resultPart2 := part2(line)
	fmt.Println("invalid total part 2:", resultPart2)
}

func part1(line string) int {
	var total int
	for r := range strings.SplitSeq(line, ",") {
		bounds := strings.Split(r, "-")
		firstId, err := strconv.Atoi(bounds[0])
		if err != nil {
			fmt.Printf("failed to convert string to integer: %v", err)
		}
		lastId, err := strconv.Atoi(bounds[len(bounds)-1])
		if err != nil {
			fmt.Printf("failed to convert string to integer: %v", err)
		}

		for i := firstId; i <= lastId; i++ {
			id := strconv.Itoa(i)
			if len(id)%2 == 0 {
				seqLength := len(id) / 2
				if id[:seqLength] == id[seqLength:] {
					total += i
				}
			}
		}
	}

	return total
}

func part2(line string) int {
	var total int
	var ids []int

	for r := range strings.SplitSeq(line, ",") {
		bounds := strings.Split(r, "-")
		firstId, err := strconv.Atoi(bounds[0])
		if err != nil {
			fmt.Printf("failed to convert string to integer: %v", err)
		}
		lastId, err := strconv.Atoi(bounds[len(bounds)-1])
		if err != nil {
			fmt.Printf("failed to convert string to integer: %v", err)
		}

		for i := firstId; i <= lastId; i++ {
			ids = append(ids, i)
		}
	}

	invalidIds := make(map[int]bool)

	for _, id := range ids {
		i := strconv.Itoa(id)
		divisors := getDivisors(len(i))

		for _, d := range divisors {
			seqLength := len(i) / d
			if _, ok := invalidIds[id]; !ok {
				if hasRepeatedSequence(i, d, seqLength) {
					total += id
					invalidIds[id] = true
				}
			}
		}
	}

	return total
}

func getDivisors(n int) []int {
	var d []int

	for i := 1; i <= n; i++ {
		if n%i == 0 && i != 1 {
			d = append(d, i)
		}
	}

	return d
}

func hasRepeatedSequence(s string, d int, split int) bool {
	var prevSeq string
	totalSeq := 1

	for i := 0; i < len(s); i += split {
		end := min(i+split, len(s))
		newSeq := s[i:end]
		if prevSeq == newSeq {
			totalSeq++
		}
		prevSeq = newSeq
	}

	return totalSeq == d
}
