package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(scanner *bufio.Scanner) int {
	cur_pos := 50
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		move, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Couldn't convert string to int")
		}

		switch direction {
		case 'L':
			cur_pos = (cur_pos - move) % 100
		case 'R':
			cur_pos = (cur_pos + move) % 100
		}

		if cur_pos == 0 {
			total++
		}
	}
	return total
}

func main() {
	file, err := os.Open("./inputs/d1i1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pass := part1(scanner)
	fmt.Printf("Password is: %d\n", pass)
}
