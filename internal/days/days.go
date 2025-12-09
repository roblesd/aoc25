package days

import (
	"io"

	"github.com/roblesd/aoc25/internal/day1"
	"github.com/roblesd/aoc25/internal/day2"
	"github.com/roblesd/aoc25/internal/day3"
	"github.com/roblesd/aoc25/internal/day4"
	// "github.com/roblesd/aoc25/internal/day5"
	// "github.com/roblesd/aoc25/internal/day6"
	// "github.com/roblesd/aoc25/internal/day7"
	// "github.com/roblesd/aoc25/internal/day8"
	// "github.com/roblesd/aoc25/internal/day9"
	// "github.com/roblesd/aoc25/internal/day10"
	// "github.com/roblesd/aoc25/internal/day11"
	// "github.com/roblesd/aoc25/internal/day12"
)

type Solver struct {
	Part1 func(io.Reader) int
	Part2 func(io.Reader) int
}

var Registry = map[int]Solver{
	1: {Part1: day1.Part1, Part2: day1.Part2},
	2: {Part1: day2.Part1, Part2: day2.Part2},
	3: {Part1: day3.Part1, Part2: day3.Part2},
	4: {Part1: day4.Part1, Part2: day4.Part2},
}
