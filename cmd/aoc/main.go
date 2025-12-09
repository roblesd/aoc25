package main

import (
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

	// Run the specified part
	switch *part {
	case 1:
		solver.Part1(file)
	case 2:
		solver.Part2(file)
	default:
		fmt.Println("Invalid part selected")
		return
	}
}
