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

type State struct {
	node string
	a    bool
	b    bool
}

func Part2(r io.Reader) int {
	graph := parseInput(r)

	var search func(string, bool, bool, map[State]int) int
	search = func(s string, dac, fft bool, cache map[State]int) int {
		//check cache first
		key := State{s, dac, fft}
		if v, ok := cache[key]; ok {
			return v
		}

		//base case
		if s == "out" {
			if dac && fft {
				return 1
			} else {
				return 0
			}
		}

		//check all outputs from current device
		total := 0
		for _, device := range graph[s] {
			total += search(device, dac || (device == "dac"), fft || (device == "fft"), cache)
		}
		cache[key] = total
		return total
	}
	cache := make(map[State]int)
	paths := search("svr", false, false, cache)
	fmt.Printf("Day 11 Part 2; Paths found: %d\n", paths)
	return paths
}
