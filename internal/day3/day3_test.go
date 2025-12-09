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
