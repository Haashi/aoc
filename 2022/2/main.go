package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parseInput() []string {
	readFile, err := os.Open("input")
	if err != nil {
		log.Fatal("failed to open file: " + err.Error())
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	readFile.Close()
	return lines
}

func main() {
	input := parseInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

// A X ROCK
// B Y PAPER
// C Z SCIZ
func part1(input []string) any {
	score := 0
	for _, line := range input {
		if line == "A X" {
			score += 1 + 3
		}
		if line == "B X" {
			score += 1
		}
		if line == "C X" {
			score += 1 + 6
		}
		if line == "A Y" {
			score += 2 + 6
		}
		if line == "B Y" {
			score += 2 + 3
		}
		if line == "C Y" {
			score += 2
		}
		if line == "A Z" {
			score += 3
		}
		if line == "B Z" {
			score += 3 + 6
		}
		if line == "C Z" {
			score += 3 + 3
		}
	}
	return score
}

// A X ROCK  LOOSE
// B Y PAPER DRAW
// C Z SCIZ  WIN
func part2(input []string) any {
	score := 0
	for _, line := range input {
		if line == "A X" {
			score += 3
		}
		if line == "B X" {
			score += 1
		}
		if line == "C X" {
			score += 2
		}
		if line == "A Y" {
			score += 3 + 1
		}
		if line == "B Y" {
			score += 3 + 2
		}
		if line == "C Y" {
			score += 3 + 3
		}
		if line == "A Z" {
			score += 6 + 2
		}
		if line == "B Z" {
			score += 6 + 3
		}
		if line == "C Z" {
			score += 6 + 1
		}
	}
	return score
}
