package day9

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Point struct {
	x int
	y int
}

func (cc1 Point) area(cc2 Point) int {
	return (abs(cc1.x-cc2.x) + 1) * (abs(cc1.y-cc2.y) + 1)
}

func parseInput(r io.Reader) []Point {
	var corners []Point
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		corners = append(corners, Point{x, y})
	}
	return corners
}

type Rectangle struct {
	p1, p2 Point
}

func (rec Rectangle) area() int {
	return rec.p1.area(rec.p2)
}

func (rec Rectangle) coordsInfo() (int, int, int, int) {
	minX := min(rec.p1.x, rec.p2.x)
	maxX := max(rec.p1.x, rec.p2.x)
	minY := min(rec.p1.y, rec.p2.y)
	maxY := max(rec.p1.y, rec.p2.y)
	return minX, maxX, minY, maxY
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
	corners := parseInput(r)
	var rectangles []Rectangle
	for i := 0; i < len(corners)-1; i++ {
		for j := i + 1; j < len(corners); j++ {
			rectangles = append(rectangles, Rectangle{corners[i], corners[j]})
		}
	}
	rectangles = append(rectangles, Rectangle{corners[len(corners)-1], corners[0]})

	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].area() > rectangles[j].area()
	})

	maxArea := 0
	for _, r := range rectangles {
		minX, maxX, minY, maxY := r.coordsInfo()
		valid := true
		for i, c1 := range corners {
			c2 := corners[(i+1)%len(corners)] // to get that last edge
			// vertical perimeter edge
			if c1.x == c2.x {
				ylMin, ylMax := min(c1.y, c2.y), max(c1.y, c2.y)
				if minX < c1.x && maxX > c1.x && !(minY >= ylMax || maxY <= ylMin) {
					valid = false
					break
				}
			} else if c1.y == c2.y {
				xlMin, xlMax := min(c1.x, c2.x), max(c1.x, c2.x)
				if minY < c1.y && maxY > c1.y && !(minX >= xlMax || maxX <= xlMin) {
					valid = false
					break
				}
			} else {
				panic("Uh oh")
			}
		}
		if valid {
			maxArea = r.area()
			fmt.Printf("Day 9 Part 2; Max Area Rectangle: %d\n", maxArea)
			return maxArea // FOUND the first valid rectangle (because sorted desc)
		}
	}
	return 0
}
