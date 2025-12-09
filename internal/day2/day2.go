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

	// For odd length numbers; nvm
	// for i := 0; i < n; i++ {
	// 	if s[i] != s[0] {
	// 		return false
	// 	}
	// }
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
			// fmt.Printf("Checking %d...", i)
			if repeats(i) {
				// fmt.Print("repeater!")
				total += i
			}
			// fmt.Println()
		}

	}
	fmt.Printf("Invalid Product ID's add up to %d\n", total)
	return total
}

func Part2(r io.Reader) int {
	// scanner := bufio.NewScanner(r)
	return 0
}
