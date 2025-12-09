package days

import (
	"io"

	"github.com/roblesd/aoc25/internal/day1"
	"github.com/roblesd/aoc25/internal/day2"
)

type Solver struct {
	Part1 func(io.Reader) int
	Part2 func(io.Reader) int
}

var Registry = map[int]Solver{
	1: {Part1: day1.Part1, Part2: day1.Part2},
	2: {Part1: day2.Part1, Part2: day2.Part2},
}
