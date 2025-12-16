package day12

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

type Region struct {
	w, l int
	fits []int
}

type Present struct {
	shape []string
	size  int
}

func (p Present) print() {
	for i, s := range p.shape {
		fmt.Printf("%d: %q\n", i, s)
	}
}

func parseInput(r io.Reader) ([]Present, []Region) {
	var presents []Present
	var regions []Region
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Presents Parsing
		if strings.HasSuffix(line, ":") && unicode.IsDigit(rune(line[0])) {
			// don't need present index
			var shape []string
			size := 0
			for i := 0; i < 3; i++ {
				scanner.Scan()
				shape = append(shape, scanner.Text())
				size += strings.Count(scanner.Text(), "#")
			}
			presents = append(presents, Present{shape, size})
			continue
		}

		// Regions Parsing
		if strings.Contains(line, "x") {
			parts := strings.Split(line, ":")
			dims := strings.Split(parts[0], "x")
			w, _ := strconv.Atoi(dims[0])
			h, _ := strconv.Atoi(dims[1])

			nums := strings.Fields(parts[1])
			fits := make([]int, len(nums))
			for i, s := range nums {
				fits[i], _ = strconv.Atoi(s)
			}
			regions = append(regions, Region{w, h, fits})
		}
	}
	return presents, regions
}

func sumSlice(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func Part1(r io.Reader) int {
	presents, regions := parseInput(r)
	total := 0
	check := make([]int, 0)
	for i, r := range regions {
		minSpace := 0
		for j, freq := range r.fits {
			minSpace += presents[j].size * freq
		}
		upperBound := (r.l / 3) * (r.w / 3)
		if minSpace > r.l*r.w {
			continue
		} else if upperBound >= sumSlice(r.fits) {
			total++
		} else {
			check = append(check, i)
		}
	}
	fmt.Printf("May want to check these: %v\n", check)
	fmt.Printf("Day 12, Part 1; Valid Regions: %d\n", total)
	return total
}

func Part2(r io.Reader) int {
	return 0
}
