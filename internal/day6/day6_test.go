package day6

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   + \n"
	out := Part1(strings.NewReader(input))
	answer := 4277556
	if out != answer {
		t.Fatalf("Part1(%q) = %d; want %d", input, out, answer)
	}
}

func TestPart2(t *testing.T) {
	input := "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   + \n"
	out := Part2(strings.NewReader(input))
	answer := 3263827
	if out != answer {
		t.Fatalf("Part2(%q) = %d; want %d", input, out, answer)
	}
}
