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

type pos struct {
	x int
	y int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func (p pos) distance(p1 pos) int {
	return max(abs(p.x-p1.x), abs(p.y-p1.y))
}

func part1(input []string) any {
	head := pos{}
	tail := pos{}
	seen := make(map[pos]bool)
	seen[pos{tail.x, tail.y}] = true
	for _, line := range input {
		dir := strings.Split(line, " ")[0]
		length, _ := strconv.Atoi(strings.Split(line, " ")[1])
		for i := 0; i < length; i++ {
			if dir == "R" {
				head.x += 1
			}
			if dir == "L" {
				head.x -= 1
			}
			if dir == "U" {
				head.y -= 1
			}
			if dir == "D" {
				head.y += 1
			}
			if tail.distance(head) > 1 {
				dx := head.x - tail.x
				dy := head.y - tail.y
				if dx > 1 {
					dx = 1
				}
				if dx < -1 {
					dx = -1
				}
				if dy > 1 {
					dy = 1
				}
				if dy < -1 {
					dy = -1
				}
				tail.x += dx
				tail.y += dy
				seen[pos{tail.x, tail.y}] = true
			}
		}
	}
	return len(seen)
}

func part2(input []string) any {
	knots := make([]pos, 10)
	seen := make(map[pos]bool)
	seen[pos{knots[9].x, knots[9].y}] = true
	for _, line := range input {
		dir := strings.Split(line, " ")[0]
		length, _ := strconv.Atoi(strings.Split(line, " ")[1])
		for i := 0; i < length; i++ {
			if dir == "R" {
				knots[0].x += 1
			}
			if dir == "L" {
				knots[0].x -= 1
			}
			if dir == "U" {
				knots[0].y -= 1
			}
			if dir == "D" {
				knots[0].y += 1
			}
			for i := 1; i < 10; i++ {
				if knots[i].distance(knots[i-1]) > 1 {
					dx := knots[i-1].x - knots[i].x
					dy := knots[i-1].y - knots[i].y
					if dx > 1 {
						dx = 1
					}
					if dx < -1 {
						dx = -1
					}
					if dy > 1 {
						dy = 1
					}
					if dy < -1 {
						dy = -1
					}
					knots[i].x += dx
					knots[i].y += dy
					seen[pos{knots[9].x, knots[9].y}] = true
				}
			}
		}
	}
	return len(seen)
}
