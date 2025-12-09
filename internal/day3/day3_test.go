package day3

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := "987654321111111\n811111111111119\n234234234234278\n818181911112111\n"
	out := Part1(strings.NewReader(input))
	answer := 357
	if out != answer {
		t.Fatalf("Part1(%q) = %d; want %d", input, out, answer)
	}
}

func TestPart2_1(t *testing.T) {
	input := "234234234234278\n"
	out := Part2(strings.NewReader(input))
	answer := 434234234278
	if out != answer {
		t.Fatalf("Part1(%q) = %d; want %d", input, out, answer)
	}
}

func TestPart2_2(t *testing.T) {
	input := "987654321111111\n811111111111119\n234234234234278\n818181911112111\n"
	out := Part2(strings.NewReader(input))
	answer := 3121910778619
	if out != answer {
		t.Fatalf("Part1(%q) = %d; want %d", input, out, answer)
	}
}
