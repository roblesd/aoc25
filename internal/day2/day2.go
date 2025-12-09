package day2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func repeats(num int) bool {
	s := strconv.Itoa(num)
	n := len(s)
	if num < 10 {
		return false
	}
	if n%2 == 0 {
		return s[:n/2] == s[n/2:]
	}

	return false
}

func repeatsMultiple(num int) bool {
	s := strconv.Itoa(num)
	n := len(s)
	if num < 10 {
		return false
	}

	// Check each substring to see if it repeats
	// can skip over substrings that dont divide the string length evenly
	// only need to go halfway because after that repeating isn't possible
	for i := 1; i <= n/2; i++ {
		if n%(i) != 0 {
			continue
		}
		curSubstring := s[:i]
		for j := i; j <= n; j += i {
			if j == n {
				return true
			}
			if s[j:j+i] != curSubstring {
				break
			}
		}
	}
	return false
}

func Part1(r io.Reader) int {
	// for some reason all the input is on one line
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	ranges := strings.Split(scanner.Text(), ",")

	total := 0
	for _, str := range ranges {
		nums := strings.Split(str, "-")
		lower, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		upper, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		for i := lower; i <= upper; i++ {
			if repeats(i) {
				total += i
			}
		}
	}
	fmt.Printf("Invalid Product ID's add up to %d\n", total)
	return total
}

func Part2(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	ranges := strings.Split(scanner.Text(), ",")

	total := 0
	for _, str := range ranges {
		nums := strings.Split(str, "-")
		lower, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		upper, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		for i := lower; i <= upper; i++ {
			if repeatsMultiple(i) {
				total += i
			}
		}
	}
	fmt.Printf("Invalid Product ID's add up to %d\n", total)
	return total
}
