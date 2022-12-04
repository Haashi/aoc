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
	ret := 0
	for _, line := range input {
		runes := []rune(line)
		runesInFirst := [58]int{}
		runesInSecond := [58]int{}
		for i := 0; i < len(runes)/2; i++ {
			r := runes[i]
			fmt.Print(string(r))
			runesInFirst[r-65]++
		}
		for i := len(runes) / 2; i < len(runes); i++ {
			r := runes[i]
			fmt.Print(string(r))
			runesInSecond[r-65]++
		}
		for i := range runesInFirst {
			if runesInFirst[i]*runesInSecond[i] > 0 {
				fmt.Println(string(rune(i + 65)))
				var res = 0
				if i > 26 {
					res = (i - 31)
				} else {
					res = i + 26 + 1
				}
				ret += res
			}
		}
	}
	return ret
}

func part2(input []string) any {
	ret := 0
	for j := 0; j < len(input); j += 3 {
		firstHalf := []rune(input[j])
		secondHalf := []rune(input[j+1])
		thirdHalf := []rune(input[j+2])
		runesInFirst := [58]int{}
		runesInSecond := [58]int{}
		runesInThird := [58]int{}
		for i := 0; i < len(firstHalf); i++ {
			r := firstHalf[i]
			runesInFirst[r-65]++
		}
		for i := 0; i < len(secondHalf); i++ {
			r := secondHalf[i]
			runesInSecond[r-65]++
		}
		for i := 0; i < len(thirdHalf); i++ {
			r := thirdHalf[i]
			runesInThird[r-65]++
		}
		for i := range runesInFirst {
			if runesInFirst[i]*runesInSecond[i]*runesInThird[i] > 0 {
				var res = 0
				if i > 26 {
					res = (i - 31)
				} else {
					res = i + 26 + 1
				}
				ret += res
				break
			}
		}
	}
	return ret
}
