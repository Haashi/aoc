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

func part1(input []string) any {
	size := len(input)
	grid := make([]int, size*size)
	visible := make(map[int]bool)
	for i, line := range input {
		str := strings.Split(line, "")
		for j := range str {
			tree, _ := strconv.Atoi(str[j])
			grid[i*size+j] = tree
		}
	}
	//look from left
	for i := 0; i < size; i++ {
		max := -1
		for j := 0; j < size; j++ {
			if grid[i*size+j] > max {
				visible[i*size+j] = true
				max = grid[i*size+j]
			}
		}
	}
	//look from right
	for i := 0; i < size; i++ {
		max := -1
		for j := size - 1; j >= 0; j-- {
			if grid[i*size+j] > max {
				visible[i*size+j] = true
				max = grid[i*size+j]
			}
		}
	}
	//look from top
	for j := 0; j < size; j++ {
		max := -1
		for i := 0; i < size; i++ {
			if grid[i*size+j] > max {
				visible[i*size+j] = true
				max = grid[i*size+j]
			}
		}
	}
	//look from bottom
	for j := 0; j < size; j++ {
		max := -1
		for i := size - 1; i >= 0; i-- {
			if grid[i*size+j] > max {
				visible[i*size+j] = true
				max = grid[i*size+j]
			}
		}
	}
	return len(visible)
}

func part2(input []string) any {
	size := len(input)
	grid := make([]int, size*size)
	for i, line := range input {
		str := strings.Split(line, "")
		for j := range str {
			tree, _ := strconv.Atoi(str[j])
			grid[i*size+j] = tree
		}
	}

	maxScenicScore := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			scenicScore := 0
			height := grid[i*size+j]
			//look left
			scenicScoreLeft := 0
			for jj := j - 1; jj >= 0; jj-- {
				tree := grid[i*size+jj]
				if tree >= height {
					scenicScoreLeft++
					break
				}
				scenicScoreLeft++
			}
			//look right
			scenicScoreRight := 0
			for jj := j + 1; jj < size; jj++ {
				tree := grid[i*size+jj]
				if tree >= height {
					scenicScoreRight++
					break
				}
				scenicScoreRight++
			}
			//look top
			scenicScoreTop := 0
			for ii := i - 1; ii >= 0; ii-- {
				tree := grid[ii*size+j]
				if tree >= height {
					scenicScoreTop++
					break
				}
				scenicScoreTop++
			}
			//look bottom
			scenicScoreBottom := 0
			for ii := i + 1; ii < size; ii++ {
				tree := grid[ii*size+j]
				if tree >= height {
					scenicScoreBottom++
					break
				}
				scenicScoreBottom++
			}

			scenicScore = scenicScoreLeft * scenicScoreRight * scenicScoreTop * scenicScoreBottom
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	return maxScenicScore
}
