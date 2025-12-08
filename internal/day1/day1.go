package day1

import (
	"bufio"
	"fmt"
	"strconv"
)

func posMod(a, n int) int {
	return (a%n + n) % n
}

func Part1(scanner *bufio.Scanner) string {
	curPos := 50
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		move, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Couldn't convert string to int")
		}

		switch direction {
		case 'L':
			curPos = (curPos - move) % 100
		case 'R':
			curPos = (curPos + move) % 100
		}

		if curPos == 0 {
			total++
		}
	}
	return fmt.Sprintf("The password is %d", total)
}

func Part2(scanner *bufio.Scanner) string {
	curPos := 50
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		move, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Couldn't convert string to int")
			continue
		}

		rotations := move / 100
		remainder := posMod(move, 100)
		total += rotations
		switch direction {
		case 'L':
			if curPos != 0 && curPos-remainder <= 0 {
				total++
			}
			// major key
			// negative values don't matter if we just care about landing on 0
			// but it messes up the calculation for moving past 0
			curPos = posMod((curPos - move), 100)
		case 'R':
			if curPos+remainder >= 100 {
				total++
			}
			curPos = (curPos + move) % 100
		}
	}
	return fmt.Sprintf("The password is %d", total)
}
