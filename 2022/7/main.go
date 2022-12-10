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

type node struct {
	size     int
	children map[string]*node
	parent   *node
	name     string
}

func part1(input []string) any {
	currentDir := &node{}
	currentDir.children = make(map[string]*node)
	rootDir := currentDir
	for i, line := range input {
		if i == 0 {
			continue
		}
		if line == "$ cd .." {
			currentDir = currentDir.parent
		} else if strings.Contains(line, "$ cd") {
			dirName := strings.Split(line, " ")[2]
			currentDir = currentDir.children[dirName]
		} else if strings.Contains(line, "dir") {
			dirName := strings.Split(line, " ")[1]
			newDir := &node{}
			newDir.children = make(map[string]*node)
			newDir.parent = currentDir
			newDir.name = dirName
			currentDir.children[dirName] = newDir
		} else if regexp.MustCompile(`\d`).MatchString(line) {
			fileName := strings.Split(line, " ")[1]
			sizeStr := strings.Split(line, " ")[0]
			fileSize, _ := strconv.ParseInt(sizeStr, 10, 64)
			newFile := &node{}
			newFile.size = int(fileSize)
			newFile.parent = currentDir
			newFile.name = fileName
			currentDir.children[fileName] = newFile
		}
	}
	sum := 0
	dfsCalcSize(rootDir, &sum)
	return sum
}

func dfsCalcSize(n *node, sum *int) {
	if n.children == nil {
		return
	} else {

		total := 0
		for _, child := range n.children {
			dfsCalcSize(child, sum)
			total += child.size
		}
		n.size = total
		if n.size <= 100000 {
			*sum = *sum + n.size
		}
	}
}

func part2(input []string) any {
	currentDir := &node{}
	currentDir.children = make(map[string]*node)
	rootDir := currentDir
	rootDir.name = "/"
	for i, line := range input {
		if i == 0 {
			continue
		}
		if line == "$ cd .." {
			currentDir = currentDir.parent
		} else if strings.Contains(line, "$ cd") {
			dirName := strings.Split(line, " ")[2]
			currentDir = currentDir.children[dirName]
		} else if strings.Contains(line, "dir") {
			dirName := strings.Split(line, " ")[1]
			newDir := &node{}
			newDir.children = make(map[string]*node)
			newDir.parent = currentDir
			newDir.name = dirName
			currentDir.children[dirName] = newDir
		} else if regexp.MustCompile(`\d`).MatchString(line) {
			fileName := strings.Split(line, " ")[1]
			sizeStr := strings.Split(line, " ")[0]
			fileSize, _ := strconv.ParseInt(sizeStr, 10, 64)
			newFile := &node{}
			newFile.size = int(fileSize)
			newFile.parent = currentDir
			newFile.name = fileName
			currentDir.children[fileName] = newFile
		}
	}
	sum := 0
	dfsCalcSize(rootDir, &sum)
	min := dfsChooseMin(rootDir, rootDir, 30000000-(70000000-rootDir.size))
	return min.size
}

func dfsChooseMin(n *node, min *node, totalSpaceNeeded int) *node {
	if n.children == nil {
		return min
	} else {
		for _, child := range n.children {
			min = dfsChooseMin(child, min, totalSpaceNeeded)
		}
		if n.size < min.size && n.size >= totalSpaceNeeded {
			return n
		}
		return min
	}
}
