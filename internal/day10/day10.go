package day10

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	lights  int
	buttons []int
	joltage []int
}

func parseLightString(lightString string) int {
	res := 0b0
	for i, c := range lightString {
		if c == '#' {
			res |= 1 << i
		}
	}
	return res
}

func parseButtonsString(buttonsString string) []int {
	tupleRe := regexp.MustCompile(`\(([^)]*)\)`)
	var buttons []int
	for _, m := range tupleRe.FindAllStringSubmatch(buttonsString, -1) {
		res := 0
		parts := strings.Split(m[1], ",")
		for _, p := range parts {
			n, _ := strconv.Atoi(p)
			res |= 1 << n
		}
		buttons = append(buttons, res)
	}
	return buttons
}

func parseJoltage(joltageString string) []int {
	var joltages []int
	for _, p := range strings.Split(joltageString, ",") {
		n, _ := strconv.Atoi(p)
		joltages = append(joltages, n)
	}
	return joltages
}

func parseInput(r io.Reader) []Machine {
	scanner := bufio.NewScanner(r)
	re := regexp.MustCompile(`\[(.*?)\]\s*((?:\([^\)]*\)\s*)+)\{(.*?)\}`)
	var machines []Machine
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		lightString := matches[1]   // ".###.#"
		buttonsString := matches[2] // "(0,1,2,3,4) (0,3,4) ..."
		joltageString := matches[3] // "10,11,11,5,10,5"

		lightNum := parseLightString(lightString)
		buttons := parseButtonsString(buttonsString)
		joltages := parseJoltage(joltageString)

		machines = append(machines, Machine{lightNum, buttons, joltages})
	}
	return machines
}

type state struct {
	presses, val int
}

func findMinPresses(m Machine) int {
	var queue []state
	seen := make(map[int]struct{})
	queue = append(queue, state{0, 0})
	head := 0
	for head < len(queue) {
		cur := queue[head]
		head++

		if cur.val == m.lights {
			return cur.presses
		}
		for _, b := range m.buttons {
			newVal := cur.val ^ b
			if _, exists := seen[newVal]; !exists {
				// if we haven't seen the value before, we'll entertain it
				queue = append(queue, state{cur.presses + 1, newVal})
				seen[newVal] = struct{}{}
			}
		}
	}
	return -1
}

func Part1(r io.Reader) int {
	machines := parseInput(r)
	total := 0
	for _, m := range machines {
		total += findMinPresses(m)
	}
	fmt.Printf("Day 10 Part 1; Minimum Presses: %d\n", total)
	return total
}

func lightPatterns(buttons []int) map[int][]int {
	patterns := make(map[int][]int)
	// 1. find all 2^n subsets of buttons
	n := len(buttons)
	total := 1 << n
	for mask := 0; mask < total; mask++ {
		lights := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				// buttons[i] is ON in this combination
				lights ^= buttons[i]
			}
		}
		patterns[lights] = append(patterns[lights], mask)
	}
	// patterns maps light patterns (as bits) to button combos that produce it
	return patterns
}

func findParity(joltages []int) int {
	// return bit representation of which lights are on
	lights := 0
	for i, joltage := range joltages {
		if joltage%2 == 1 {
			lights |= 1 << i
		}
	}
	return lights
}

func allZero(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

func encode(target []int) string {
	var b strings.Builder
	b.Grow(len(target) * 3) // rough guess
	for _, v := range target {
		b.WriteString(strconv.Itoa(v))
		b.WriteByte(',')
	}
	return b.String()
}

func configureJoltages(buttons []int, joltages []int, patterns map[int][]int) int {
	cache := make(map[string]int)
	numButtons := len(buttons)
	var minPresses func(target []int) int
	minPresses = func(target []int) int {
		key := encode(target)
		// lookup
		if v, ok := cache[key]; ok {
			return v
		}
		// base case, no presses needed at 0'd out joltage target
		if allZero(target) {
			cache[key] = 0
			return 0
		}
		lights := findParity(target)
		newTarget := make([]int, len(target))
		answer := 1000000 //"infinity"
		for _, buttonMask := range patterns[lights] {
			copy(newTarget, target)
			presses := 0
			for i := 0; i < numButtons; i++ {
				// determine which buttons are 'on'
				// update the target joltage
				if buttonMask&(1<<i) != 0 {
					presses++
					//need to index into buttons
					joltageMask := buttons[i]
					for j := 0; joltageMask > 0; j++ {
						if joltageMask&1 == 1 {
							// bit j is ON
							newTarget[j]--
						}
						joltageMask >>= 1
					}
				}
			}
			// check that joltages are non-negative
			valid := true
			for _, n := range newTarget {
				if n < 0 {
					valid = false
				}
			}
			if !valid {
				continue
			}

			// Now, all target joltages should be even, so we can split
			halfTarget := make([]int, len(target))
			for i, joltage := range newTarget {
				halfTarget[i] = joltage / 2
			}
			answer = min(answer, presses+2*minPresses(halfTarget))
		}
		cache[key] = answer
		return answer
	}
	return minPresses(joltages)
}

func Part2(r io.Reader) int {
	machines := parseInput(r)
	total := 0
	for _, m := range machines {
		patterns := lightPatterns(m.buttons)
		ans := configureJoltages(m.buttons, m.joltage, patterns)
		total += ans
	}
	fmt.Printf("Day 10 Part 2; Joltage configured in %d presses\n", total)
	return total
}
