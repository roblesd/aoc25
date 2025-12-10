package day8

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}

type PairDist struct {
	distsq int
	i      int
	j      int
}

func parsePoints(r io.Reader) []Point {
	var points []Point
	scanner := bufio.NewScanner(r)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		z, _ := strconv.Atoi(nums[2])
		points = append(points, Point{x, y, z})
		i++
	}
	return points
}

func findDistances(points []Point) []PairDist {
	var distances []PairDist
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			dist := (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y) + (p1.z-p2.z)*(p1.z-p2.z)
			distances = append(distances, PairDist{dist, i, j})
		}
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distsq < distances[j].distsq
	})
	return distances
}

func joinBoxes(points []Point, distances []PairDist, limit int) int {
	// Union-find approach
	// each point starts out as their own, and each has a size of 1
	parents := make([]int, len(points))
	for i := range parents {
		parents[i] = i
	}
	sizes := make([]int, len(points))
	for i := range sizes {
		sizes[i] = 1
	}

	clusters := len(points)

	var find func(int) int
	find = func(i int) int {
		if parents[i] == i {
			return i
		}
		parents[i] = find(parents[i])
		return parents[i]
	}

	union := func(i int, j int) {
		iParent, jParent := find(i), find(j)
		if iParent == jParent {
			return
		}
		if sizes[iParent] > sizes[jParent] {
			parents[jParent] = iParent
			sizes[iParent] += sizes[jParent]
		} else {
			parents[iParent] = jParent
			sizes[jParent] += sizes[iParent]
		}
		clusters--
	}

	for i := 0; i < limit; i++ {
		distPair := distances[i]
		union(distPair.i, distPair.j)
		if clusters == 1 {
			return points[distPair.i].x * points[distPair.j].x
		}
	}

	sort.Ints(sizes)
	total := sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
	return total
}

func Part1(r io.Reader) int {
	points := parsePoints(r)
	distances := findDistances(points)
	total := joinBoxes(points, distances, 1000)
	fmt.Printf("Day 8 Part 1; Result: %d\n", total)
	return total
}

func Part2(r io.Reader) int {
	points := parsePoints(r)
	distances := findDistances(points)
	total := joinBoxes(points, distances, len(distances))
	fmt.Printf("Day 8 Part 2; Result: %d\n", total)
	return total
}
