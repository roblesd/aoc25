package day2

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := "11-22,95-115\n"
	out := Part1(strings.NewReader(input))
	answer := 132
	if out != answer {
		t.Fatalf("Part1(%q) = %d; want %d", input, out, answer)
	}
}
