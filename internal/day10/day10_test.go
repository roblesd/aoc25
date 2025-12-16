package day10

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
			strings.NewReader("[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}\n[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}\n[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}\n"),
			7,
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

func Test_findParity(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		joltages []int
		want     int
	}{
		{"basic", []int{3, 5, 4, 7}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findParity(tt.joltages)
			if got != tt.want {
				t.Errorf("findParity() = %v, want %v", got, tt.want)
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
			strings.NewReader("[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}\n"),
			10,
		},
		{
			"Less Basic",
			strings.NewReader("[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}\n[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}\n[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}\n"),
			33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.r)
			if got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
