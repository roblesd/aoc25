package day4

import (
	"bufio"
	"fmt"
	"io"
)

func readLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func Part1(r io.Reader) int {
	grid, err := readLines(r)
	if err != nil {
		panic(err)
	}

	rows, cols := len(grid), len(grid[0])
	accessibleRolls := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '@' {
				adjacentRolls := 0
				// check all adjacent spots
				// North
				if i-1 >= 0 && grid[i-1][j] == '@' {
					adjacentRolls++
				}
				// North East
				if i-1 >= 0 && j+1 < cols && grid[i-1][j+1] == '@' {
					adjacentRolls++
				}
				// East
				if j+1 < cols && grid[i][j+1] == '@' {
					adjacentRolls++
				}
				// South East
				if i+1 < rows && j+1 < cols && grid[i+1][j+1] == '@' {
					adjacentRolls++
				}
				// South
				if i+1 < rows && grid[i+1][j] == '@' {
					adjacentRolls++
				}
				// South West
				if i+1 < rows && j-1 >= 0 && grid[i+1][j-1] == '@' {
					adjacentRolls++
				}
				// West
				if j-1 >= 0 && grid[i][j-1] == '@' {
					adjacentRolls++
				}
				// North West
				if i-1 >= 0 && j-1 >= 0 && grid[i-1][j-1] == '@' {
					adjacentRolls++
				}
				if adjacentRolls < 4 {
					accessibleRolls++
				}
			}
		}
	}
	fmt.Printf("Day 4 Part 1; Accessible Rolls: %d\n", accessibleRolls)
	return accessibleRolls
}

func Part2(r io.Reader) int {
	return 0
}
