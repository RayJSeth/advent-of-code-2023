package day4

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"

	"rayjseth.io/advent2023/util"
)

func GetAnswers() {
	lines, err := util.ReadFileToLines("./day4/input")
	if err != nil {
		log.Fatal("Input file not found")
	}
	fmt.Println("Day4:")
	part1(lines)
	fmt.Print("\n")
}

func part1(lines []string) {

	total := 0
	for _, line := range lines {
		subTotal := 0
		winningNumbers, matchNumbers := parseCard(line)
		for _, winningNumber := range winningNumbers {
			if slices.Contains(matchNumbers, winningNumber) {
				if subTotal == 0 {
					subTotal = 1
				} else {
					subTotal *= 2
				}
			}
		}
		total += subTotal
	}
	fmt.Println("Part1:", total)
}

func parseCard(line string) ([]string, []string) {
	numberBreakPattern := regexp.MustCompile(`\s+`)
	ticketBreakPattern := regexp.MustCompile(`\|\s+`)

	cardContent := strings.Split(line, ": ")[1]
	cardSections := ticketBreakPattern.Split(cardContent, -1)

	winningNumbers := numberBreakPattern.Split(cardSections[0], -1)
	matchNumbers := numberBreakPattern.Split(cardSections[1], -1)

	return winningNumbers, matchNumbers
}
