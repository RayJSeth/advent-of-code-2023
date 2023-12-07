package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"rayjseth.io/advent2023/util"
)

var expected = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func GetAnswers() {
	lines, err := util.ReadFileToLines("./day2/input")
	if err != nil {
		log.Fatal("Input file not found")
	}

	fmt.Println("Day2:")
	part1(lines)
	part2(lines)
	fmt.Print("\n")
}

func part1(lines []string) {
	total := 0

	for _, line := range lines {
		total += getGameNumIfValid(line)
	}

	fmt.Println("Part1:", total)
}

func part2(lines []string) {
	total := 0

	for _, line := range lines {
		total += getPowerOfMinViableGame(line)
	}

	fmt.Println("Part2:", total)
}

func getGameNumIfValid(s string) int {
	colIdx := strings.Index(s, ":")
	gn, _ := strconv.Atoi(string([]rune(s)[5:colIdx]))

	strChunk := ""
	for i := colIdx; i < len(s); i++ {
		cs := string(s[i])
		strChunk += cs
		for k, v := range expected {
			if strings.HasSuffix(strChunk, k) {
				numColor, _ := strconv.Atoi(string([]rune(strChunk)[2 : strings.Index(strChunk, k)-1]))
				if numColor > v {
					return 0
				}
				strChunk = ""
			}
		}
	}

	return gn
}

func getPowerOfMinViableGame(s string) int {
	power := 1
	maxEncountered := map[string]int{}
	colIdx := strings.Index(s, ":")

	strChunk := ""
	for i := colIdx; i < len(s); i++ {
		cs := string(s[i])
		strChunk += cs
		for k := range expected {
			if strings.HasSuffix(strChunk, k) {
				numColor, _ := strconv.Atoi(string([]rune(strChunk)[2 : strings.Index(strChunk, k)-1]))
				_, exists := maxEncountered[k]
				if !exists {
					maxEncountered[k] = numColor
				} else if maxEncountered[k] < numColor {
					maxEncountered[k] = numColor
				}

				strChunk = ""
			}
		}
	}

	for _, v := range maxEncountered {
		power *= v
	}

	return power
}
