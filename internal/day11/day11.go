package day11

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Graph map[string][]string

func parseInput(r io.Reader) Graph {
	graph := make(Graph)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, ":")
		edges := strings.Split(strings.TrimSpace(temp[1]), " ")
		graph[temp[0]] = append([]string(nil), edges...)
	}
	return graph
}

func Part1(r io.Reader) int {
	graph := parseInput(r)

	var search func(string) int // takes in the current device name, and returns the paths to 'out'
	search = func(s string) int {
		if s == "out" {
			return 1
		}
		total := 0
		for _, device := range graph[s] {
			total += search(device)
		}
		return total
	}
	paths := search("you")
	fmt.Printf("Day 11 Part 1; Paths found: %d\n", paths)
	return paths
}

func Part2(r io.Reader) int {
	return 0
}
