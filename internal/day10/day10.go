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

func Part2(r io.Reader) int {
	return 0
}
