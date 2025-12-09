package day3

import (
	"bufio"
	"fmt"
	"io"
)

func Part1(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	totalJoltage := 0
	for scanner.Scan() {
		line := scanner.Text()

		// find the maximum value for a digit in the tens spot and where it occurs
		maxTensDigit := 0
		idx := 0
		for i := 0; i < len(line)-1; i++ {
			curDigit := int(line[i] - '0')
			if curDigit > maxTensDigit {
				maxTensDigit = curDigit
				idx = i
			}
		}

		// find the max ones spot digit, which will come after where the tens digit was found
		maxOnesDigit := 0
		for i := idx + 1; i < len(line); i++ {
			curDigit := int(line[i] - '0')
			if curDigit > maxOnesDigit {
				maxOnesDigit = curDigit
			}
		}
		totalJoltage += (maxTensDigit*10 + maxOnesDigit)
	}
	fmt.Printf("Day 3 Part 1; Total Joltage: %d\n", totalJoltage)
	return totalJoltage
}

func Part2(r io.Reader) int {
	return 0
}
