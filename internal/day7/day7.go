package day7

import (
	"bufio"
	"fmt"
	"io"
)

func Part1(r io.Reader) int {
	var diagram []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		diagram = append(diagram, line)
	}
	rows, cols := len(diagram), len(diagram[0])

	// set up a slice to track where beams are
	beams := make([]int, cols)
	for i, char := range diagram[0] {
		if char == 'S' {
			beams[i] = 1
		} else {
			beams[i] = 0
		}
	}

	total := 0
	for i := 1; i < rows; i++ {
		temp := make([]int, len(beams))
		copy(temp, beams)

		for j := 0; j < cols; j++ {
			if diagram[i][j] == '^' && beams[j] == 1 {
				total++
				if j-1 >= 0 {
					temp[j-1] = 1
				}
				if j+1 < cols {
					temp[j+1] = 1
				}
				temp[j] = 0
			}
			copy(beams, temp)
		}
		// fmt.Printf("%d: ", i)
		// fmt.Print(temp)
		// fmt.Println()

	}
	// for _, beam := range beams {
	// 	if beam == 1 {
	// 		total++
	// 	}
	// }
	fmt.Printf("Day 7 Part 1; Beam Split %d times\n", total)
	return total
}

func Part2(r io.Reader) int {
	return 0
}
