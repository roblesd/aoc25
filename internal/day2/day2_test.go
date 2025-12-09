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

func TestPart2_1(t *testing.T) {
	input := "110-115,121212-121213\n"
	out := Part2(strings.NewReader(input))
	answer := 111 + 121212
	if out != answer {
		t.Fatalf("Part2(%q) = %d; want %d", input, out, answer)
	}
}

func TestPart2_2(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124\n"
	out := Part2(strings.NewReader(input))
	answer := 4174379265
	if out != answer {
		t.Fatalf("Part2(%q) = %d; want %d", input, out, answer)
	}
}

func TestPart2_3(t *testing.T) {
	input := "1188511885-1188511886\n"
	out := Part2(strings.NewReader(input))
	answer := 1188511885
	if out != answer {
		t.Fatalf("Part2(%q) = %d; want %d", input, out, answer)
	}
}
