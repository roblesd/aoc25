package day5

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32\n"
	out := Part1(strings.NewReader(input))
	answer := 3
	if out != answer {
		t.Fatalf("Part1(%q) = %d; want %d", input, out, answer)
	}
}

func TestPart2_1(t *testing.T) {
	input := "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32\n"
	out := Part2(strings.NewReader(input))
	answer := 14
	if out != answer {
		t.Fatalf("Part2(%q) = %d; want %d", input, out, answer)
	}
}

func TestPart2_2(t *testing.T) {
	input := "3-5\n6-10\n7-10\n7-11\n\n1\n"
	out := Part2(strings.NewReader(input))
	answer := 9
	if out != answer {
		t.Fatalf("Part2(%q) = %d; want %d", input, out, answer)
	}
}
