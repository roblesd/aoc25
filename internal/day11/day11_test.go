package day11

import (
	"io"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		r    io.Reader
		want int
	}{
		{
			"Basic",
			strings.NewReader("aaa: you hhh\nyou: bbb ccc\nbbb: ddd eee\nccc: ddd eee fff\nddd: ggg\neee: out\nfff: out\nggg: out\nhhh: ccc fff iii\niii: out\n"),
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.r)
			if got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
