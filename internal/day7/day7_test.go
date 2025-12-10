package day7

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
			strings.NewReader(".......S.......\n...............\n.......^.......\n...............\n......^.^......\n...............\n.....^.^.^.....\n...............\n....^.^...^....\n...............\n...^.^...^.^...\n...............\n..^...^.....^..\n...............\n.^.^.^.^.^...^.\n...............\n"),
			21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.r)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestPart2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		r    io.Reader
		want int
	}{
		{
			"Basic",
			strings.NewReader(".......S.......\n...............\n.......^.......\n...............\n......^.^......\n...............\n.....^.^.^.....\n...............\n....^.^...^....\n...............\n...^.^...^.^...\n...............\n..^...^.....^..\n...............\n.^.^.^.^.^...^.\n...............\n"),
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.r)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
