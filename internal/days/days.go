package days

import (
	"io"

	"github.com/roblesd/aoc25/internal/day1"
	"github.com/roblesd/aoc25/internal/day10"
	"github.com/roblesd/aoc25/internal/day11"
	"github.com/roblesd/aoc25/internal/day12"
	"github.com/roblesd/aoc25/internal/day2"
	"github.com/roblesd/aoc25/internal/day3"
	"github.com/roblesd/aoc25/internal/day4"
	"github.com/roblesd/aoc25/internal/day5"
	"github.com/roblesd/aoc25/internal/day6"
	"github.com/roblesd/aoc25/internal/day7"
	"github.com/roblesd/aoc25/internal/day8"
	"github.com/roblesd/aoc25/internal/day9"
)

type Solver struct {
	Part1 func(io.Reader) int
	Part2 func(io.Reader) int
}

var Registry = map[int]Solver{
	1:  {Part1: day1.Part1, Part2: day1.Part2},
	2:  {Part1: day2.Part1, Part2: day2.Part2},
	3:  {Part1: day3.Part1, Part2: day3.Part2},
	4:  {Part1: day4.Part1, Part2: day4.Part2},
	5:  {Part1: day5.Part1, Part2: day5.Part2},
	6:  {Part1: day6.Part1, Part2: day6.Part2},
	7:  {Part1: day7.Part1, Part2: day7.Part2},
	8:  {Part1: day8.Part1, Part2: day8.Part2},
	9:  {Part1: day9.Part1, Part2: day9.Part2},
	10: {Part1: day10.Part1, Part2: day10.Part2},
	11: {Part1: day11.Part1, Part2: day11.Part2},
	12: {Part1: day12.Part1, Part2: day12.Part2},
}
