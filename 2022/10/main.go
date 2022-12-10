package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

type op struct {
	typ            string
	timeToComplete int
	amount         int
}

func part1(input []string) any {
	opes := make([]op, 0)
	for _, line := range input {
		if line == "noop" {
			opes = append(opes, op{typ: "noop", timeToComplete: 1})
		} else {
			split := strings.Split(line, " ")
			amount, _ := strconv.Atoi(split[1])
			opes = append(opes, op{typ: "addx", timeToComplete: 2, amount: amount})
		}
	}
	cycle := 0
	signal := 0
	importantCycles := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
	}
	val := 1
	for _, op := range opes {
		for op.timeToComplete > 0 {
			cycle++
			if importantCycles[cycle] {
				signal += cycle * val
			}
			op.timeToComplete--
		}
		if op.typ == "addx" {
			val += op.amount
		} else {

		}
	}
	return signal
}

func part2(input []string) any {
	opes := make([]op, 0)
	for _, line := range input {
		if line == "noop" {
			opes = append(opes, op{typ: "noop", timeToComplete: 1})
		} else {
			split := strings.Split(line, " ")
			amount, _ := strconv.Atoi(split[1])
			opes = append(opes, op{typ: "addx", timeToComplete: 2, amount: amount})
		}
	}
	cycle := 0
	val := 1
	crtScreen := make([][]rune, 6)
	for i := range crtScreen {
		crtScreen[i] = make([]rune, 40)
		for j := range crtScreen[i] {
			crtScreen[i][j] = '.'
		}
	}
	for _, op := range opes {
		for op.timeToComplete > 0 {
			if cycle%40 == val-1 {
				crtScreen[cycle/40][val-1] = '#'
			}
			if cycle%40 == val {
				crtScreen[cycle/40][val] = '#'
			}
			if cycle%40 == val+1 {
				crtScreen[cycle/40][val+1] = '#'
			}
			cycle++
			op.timeToComplete--
		}
		if op.typ == "addx" {
			val += op.amount
		} else {

		}
	}
	crtScreenToString := "\n"
	for _, line := range crtScreen {
		crtScreenToString += string(line) + "\n"
	}
	return crtScreenToString
}
