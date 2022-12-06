package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

type stack struct {
	items []string
}

func (s *stack) pop() string {
	ret := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return ret
}

func (s *stack) popAmount(amount int) []string {
	ret := s.items[len(s.items)-amount:]
	s.items = s.items[:len(s.items)-amount]
	return ret
}

func (s *stack) add(item string) {
	s.items = append(s.items, item)
}

func (s *stack) addAmount(item []string) {
	s.items = append(s.items, item...)
}

func (s *stack) reverse() {
	for i, j := 0, len(s.items)-1; i < j; i, j = i+1, j-1 {
		s.items[i], s.items[j] = s.items[j], s.items[i]
	}
}

func part1(input []string) any {
	stacks := make([]stack, 0)
	for _, line := range input {
		if !strings.Contains(line, "move") && !strings.Contains(line, "1") {
			for i := 1; i < len(line); i += 4 {
				stacksIdx := i / 4
				if len(stacks) <= stacksIdx {
					stack := stack{}
					stack.items = make([]string, 0)
					stacks = append(stacks, stack)
				}
				if string(line[i]) != " " {
					stacks[stacksIdx].reverse()
					stacks[stacksIdx].add(string(line[i]))
					stacks[stacksIdx].reverse()
				}
			}
		} else if strings.Contains(line, "move") {
			r := regexp.MustCompile("[0-9]+")
			numbers := r.FindAllString(line, -1)
			amount, _ := strconv.ParseInt(numbers[0], 10, 64)
			src, _ := strconv.ParseInt(numbers[1], 10, 64)
			dest, _ := strconv.ParseInt(numbers[2], 10, 64)
			src -= 1
			dest -= 1
			for i := 0; i < int(amount); i++ {
				box := stacks[src].pop()
				stacks[dest].add(box)
			}
		}
	}
	ret := ""
	for i := 0; i < len(stacks); i++ {
		ret += stacks[i].pop()
	}
	return ret
}

func part2(input []string) any {
	stacks := make([]stack, 0)
	for _, line := range input {
		if !strings.Contains(line, "move") && !strings.Contains(line, "1") {
			for i := 1; i < len(line); i += 4 {
				stacksIdx := i / 4
				if len(stacks) <= stacksIdx {
					stack := stack{}
					stack.items = make([]string, 0)
					stacks = append(stacks, stack)
				}
				if string(line[i]) != " " {
					stacks[stacksIdx].reverse()
					stacks[stacksIdx].add(string(line[i]))
					stacks[stacksIdx].reverse()
				}
			}
		} else if strings.Contains(line, "move") {
			r := regexp.MustCompile("[0-9]+")
			numbers := r.FindAllString(line, -1)
			amount, _ := strconv.ParseInt(numbers[0], 10, 64)
			src, _ := strconv.ParseInt(numbers[1], 10, 64)
			dest, _ := strconv.ParseInt(numbers[2], 10, 64)
			src -= 1
			dest -= 1
			box := stacks[src].popAmount(int(amount))
			stacks[dest].addAmount(box)
		}
	}
	ret := ""
	for i := 0; i < len(stacks); i++ {
		ret += stacks[i].pop()
	}
	return ret
}
