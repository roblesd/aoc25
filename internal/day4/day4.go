package day4

import (
	"bufio"
	"fmt"
	"io"
)

func readLines(r io.Reader) [][]byte {
	scanner := bufio.NewScanner(r)

	var grid [][]byte

	for scanner.Scan() {
		line := scanner.Bytes()
		// MUST copy, because scanner.Bytes() is reused
		b := append([]byte(nil), line...)
		grid = append(grid, b)
	}
	return grid
}

func copyGrid(grid [][]byte) [][]byte {
	newGrid := make([][]byte, len(grid))
	for i := range grid {
		// copy each row
		newRow := make([]byte, len(grid[i]))
		copy(newRow, grid[i])
		newGrid[i] = newRow
	}
	return newGrid
}

func processGrid(grid [][]byte) ([][]byte, int) {
	markedGrid := copyGrid(grid)
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
					markedGrid[i][j] = 'X'
					accessibleRolls++
				}
			}
		}
	}
	return markedGrid, accessibleRolls
}

func Part1(r io.Reader) int {
	grid := readLines(r)
	_, accessibleRolls := processGrid(grid)
	fmt.Printf("Day 4 Part 1; Accessible Rolls: %d\n", accessibleRolls)
	return accessibleRolls
}

func Part2(r io.Reader) int {
	grid := readLines(r)
	newGrid, accessibleRolls := processGrid(grid)
	total := accessibleRolls

	for accessibleRolls > 0 {
		newGrid, accessibleRolls = processGrid(newGrid)
		total += accessibleRolls
	}
	fmt.Printf("Day 4 Part 2; Accessible Rolls: %d\n", total)
	return total
}
