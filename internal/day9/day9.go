package day9

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type CarpetCorner struct {
	x int
	y int
}

func (cc1 CarpetCorner) area(cc2 CarpetCorner) int {
	return (abs(cc1.x-cc2.x) + 1) * (abs(cc1.y-cc2.y) + 1)
}

func parseInput(r io.Reader) []CarpetCorner {
	var corners []CarpetCorner
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		corners = append(corners, CarpetCorner{x, y})
	}
	return corners
}

func Part1(r io.Reader) int {
	corners := parseInput(r)
	maxArea := 0
	for i := 0; i < len(corners); i++ {
		for j := i + 1; j < len(corners); j++ {
			area := corners[i].area(corners[j])
			maxArea = max(maxArea, area)
		}
	}
	fmt.Printf("Day 9 Part 1; Max red carpet area: %d\n", maxArea)
	return maxArea
}

func Part2(r io.Reader) int {
	return 0
}
