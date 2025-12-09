package day4

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := "..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@.\n"
	out := Part1(strings.NewReader(input))
	answer := 13
	if out != answer {
		t.Fatalf("Part1(%q) = %d; want %d", input, out, answer)
	}
}

func TestPart2(t *testing.T) {
	input := "..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@.\n"
	out := Part2(strings.NewReader(input))
	answer := 43
	if out != answer {
		t.Fatalf("Part2(%q) = %d; want %d", input, out, answer)
	}
}
