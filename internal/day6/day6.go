package day6

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func parseInput(r io.Reader) ([][]int, []string) {
	scanner := bufio.NewScanner(r)
	var worksheet [][]int
	var ops []string
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '*' || line[0] == '+' {
			ops = strings.Fields(line)
			continue
		}
		nums := strings.Fields(line)
		temp := []int{}
		for _, str := range nums {
			x, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			temp = append(temp, x)
		}
		worksheet = append(worksheet, temp)
	}
	return worksheet, ops
}

func Part1(r io.Reader) int {
	worksheet, ops := parseInput(r)
	rows, cols := len(worksheet), len(worksheet[0])
	total := 0
	for j := 0; j < cols; j++ {
		op := ops[j]
		sum := 0
		if op == "*" {
			sum = 1
		}

		for i := 0; i < rows; i++ {
			switch op {
			case "*":
				sum *= worksheet[i][j]
			case "+":
				sum += worksheet[i][j]
			}
		}
		total += sum
	}
	// fmt.Println(worksheet)
	// fmt.Println(ops)
	fmt.Printf("Day 6 Part 1; Grand Total is: %d\n", total)
	return total
}

func parseInputP2(r io.Reader) ([][]rune, []string) {
	var rawWorksheet [][]rune
	var rawOps []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		if runes[0] == '*' || runes[0] == '+' {
			rawOps = strings.Fields(scanner.Text())
		} else {
			rawWorksheet = append(rawWorksheet, runes)
		}
	}
	return rawWorksheet, rawOps
}

func transpose(ws [][]rune) [][]rune {
	rows, cols := len(ws), len(ws[0])

	out := make([][]rune, cols)
	for i := range out {
		out[i] = make([]rune, rows)
	}

	for r := range ws {
		for c := range ws[r] {
			out[c][r] = ws[r][c]
		}
	}
	return out
}

func Part2(r io.Reader) int {
	rawWorksheet, ops := parseInputP2(r)
	wsTransposed := transpose(rawWorksheet)
	rows, _ := len(wsTransposed), len(wsTransposed[0])
	total, opIdx, sum := 0, 0, 0

	for i := 0; i < rows; i++ {
		curOp := ops[opIdx]
		if curOp == "*" && sum == 0 {
			sum = 1
		}

		s := strings.TrimSpace(string(wsTransposed[i]))
		if s != "" {
			n, _ := strconv.Atoi(s)
			switch curOp {
			case "*":
				sum *= n
			case "+":
				sum += n
			}
		} else {
			total += sum
			sum = 0
			opIdx++
		}
	}
	total += sum
	fmt.Printf("Day 6 Part 2; Grand Total is: %d\n", total)
	return total
}
