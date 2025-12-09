package day5

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type IDRange struct {
	// store upper and lower bounds for the ingredient ID range
	lb int
	ub int
}

func parseInput(r io.Reader) ([]IDRange, []int) {
	var ranges []IDRange
	var ids []int
	mode := 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		// A blank new line separates the fresh ingredient ranges from the ingredient IDs
		if line == "" {
			mode = 1
			continue
		}
		switch mode {
		case 0:
			temp := strings.Split(line, "-")
			lower, err := strconv.Atoi(temp[0])
			if err != nil {
				panic(err)
			}
			upper, err := strconv.Atoi(temp[1])
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, IDRange{lower, upper})
		case 1:
			id, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			ids = append(ids, id)
		}
	}
	return ranges, ids

}
func Part1(r io.Reader) int {
	freshRanges, ids := parseInput(r)
	numFresh := 0
	for _, id := range ids {
		for _, freshRange := range freshRanges {
			if id >= freshRange.lb && id <= freshRange.ub {
				numFresh++
				break
			}
		}
	}
	fmt.Printf("Day 5 Part 1; Number of fresh ingredients: %d\n", numFresh)
	return numFresh
}

func Part2(r io.Reader) int {
	freshRanges, _ := parseInput(r)
	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i].lb < freshRanges[j].lb
	})

	var mergedRanges []IDRange
	for _, fr := range freshRanges {
		if len(mergedRanges) > 0 {
			lastIdx := len(mergedRanges) - 1
			prevRange := mergedRanges[lastIdx]
			// self contained check
			if fr.lb >= prevRange.lb && fr.ub <= prevRange.ub {
				continue
			}
			// overlap
			if fr.lb <= (prevRange.ub + 1) {
				// can do because we already checked if the new interval fits into the previous
				mergedRanges[lastIdx].ub = fr.ub
				continue
			}
		}
		//no overlap, just add the new range to merged
		mergedRanges = append(mergedRanges, fr)
	}

	total := 0
	for _, mr := range mergedRanges {
		total += (mr.ub - mr.lb + 1)
	}
	// fmt.Println(mergedRanges)
	fmt.Printf("Day 5 Part 2; Number of fresh ingredient IDs: %d\n", total)
	return total
}
