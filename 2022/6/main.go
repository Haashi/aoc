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
	fmt.Println("part1 result:", part1(input))
	fmt.Println("part2 result:", part2(input))
}

func part1(input []string) any {
	size := 4
	line := input[0]
	for i := size - 1; i < len(line); i++ {
		set := make(map[byte]int)
		for j := 0; j < size; j++ {
			k := i - j
			set[line[k]]++
		}
		if len(set) == size {
			return i + 1
		}
	}
	return 0
}

func part2(input []string) any {
	size := 14
	line := input[0]
	for i := size - 1; i < len(line); i++ {
		set := make(map[byte]int)
		for j := 0; j < size; j++ {
			k := i - j
			set[line[k]]++
		}
		if len(set) == size {
			return i + 1
		}
	}
	return 0
}
