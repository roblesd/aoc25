package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/roblesd/aoc25/internal/days"
)

func main() {
	// Set up flags
	day := flag.Int("day", 1, "Which day to run")
	part := flag.Int("part", 1, "Part 1 or 2")
	input := flag.String("input", "./inputs/day1.txt", "Input file")
	flag.Parse()

	// Grab the solver for the day specified
	solver, ok := days.Registry[*day]
	if !ok {
		fmt.Println("Unknown day")
		return
	}

	// Read in the input file
	file, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Run the specified part
	switch *part {
	case 1:
		fmt.Println(solver.Part1(scanner))
	case 2:
		fmt.Println(solver.Part2(scanner))
	default:
		fmt.Println("Invalid part selected")
		return
	}
}
