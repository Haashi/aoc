package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	ret := 0
	for _, line := range input {
		r := regexp.MustCompile("[0-9]+")
		numbers := r.FindAllString(line, -1)
		number1, _ := strconv.ParseInt(numbers[0], 10, 64)
		number2, _ := strconv.ParseInt(numbers[1], 10, 64)
		number3, _ := strconv.ParseInt(numbers[2], 10, 64)
		number4, _ := strconv.ParseInt(numbers[3], 10, 64)
		if (number1-number3)*(number2-number4) <= 0 {
			ret++
		}
	}
	return ret
}

func part2(input []string) any {
	ret := 0
	for _, line := range input {
		r := regexp.MustCompile("[0-9]+")
		numbers := r.FindAllString(line, -1)
		number1, _ := strconv.ParseInt(numbers[0], 10, 64)
		number2, _ := strconv.ParseInt(numbers[1], 10, 64)
		number3, _ := strconv.ParseInt(numbers[2], 10, 64)
		number4, _ := strconv.ParseInt(numbers[3], 10, 64)
		if number1 <= number4 && number3 <= number2 {
			ret++
		}
	}
	return ret
}
