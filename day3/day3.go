package day3

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	"rayjseth.io/advent2023/util"
)

type part struct {
	number int
	coords []int
}

func GetAnswers() {
	lines, err := util.ReadFileToLines("./day3/input")
	if err != nil {
		log.Fatal("Input file not found")
	}
	fmt.Println("Day3:")
	part1(lines)

	fmt.Print("\n")
}

func part1(lines []string) {
	total := 0

	prevLine := lines[0]
	prevLineParts := mapLineParts(prevLine)
	prevLineSymbolCoords := plotLineSymbols(prevLine)

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		lineParts := mapLineParts(line)
		lineSymbolCoords := plotLineSymbols(line)

		// anchor on previous line to look for symbols downwards or diagonally down
		for _, prevLinePartNumber := range prevLineParts {
			for _, idx := range prevLinePartNumber.coords {
				isLeftmostIdx := slices.Index(prevLinePartNumber.coords, idx) == 0
				isRightmostIdx := prevLinePartNumber.coords[len(prevLinePartNumber.coords)-1] == idx
				isOnLastLine := !(i < len(lines))
				isFirstRune := idx == 0
				isLastRune := idx == len(line)

				// only look down if there is a row below
				if !isOnLastLine {
					// look straight down
					if slices.Contains(lineSymbolCoords, idx) {
						total += prevLinePartNumber.number
					}
					// only look downleft if there is a column left and this is the leftmost rune of a part sequence
					if !isFirstRune && isLeftmostIdx {
						// look downleft
						if slices.Contains(lineSymbolCoords, idx-1) {
							total += prevLinePartNumber.number
						}
					}
					// only look downright if there is a column right and this is the rightmost rune of a part sequence
					if !isLastRune && isRightmostIdx {
						// look downright
						if slices.Contains(lineSymbolCoords, idx+1) {
							total += prevLinePartNumber.number
						}
					}
				}

			}
		}

		// anchor on current line to look for symbols upwards, diagonally upwards, or laterally
		for _, linePartNumber := range lineParts {
			for _, idx := range linePartNumber.coords {
				isLeftmostIdx := slices.Index(linePartNumber.coords, idx) == 0
				isRightmostIdx := linePartNumber.coords[len(linePartNumber.coords)-1] == idx
				isFirstRune := idx == 0
				isLastRune := idx == len(line)

				// always look up since start on 1th line
				if slices.Contains(prevLineSymbolCoords, idx) {
					total += linePartNumber.number
				}

				// only look upleft + left if there is a column left and this is the leftmost rune of a part sequence
				if !isFirstRune && isLeftmostIdx {
					// look upleft and left
					if slices.Contains(prevLineSymbolCoords, idx-1) || slices.Contains(lineSymbolCoords, idx-1) {
						total += linePartNumber.number
					}
				}
				// only look upright + right if there is a column right and this is the rightmost rune of a part sequence
				if !isLastRune && isRightmostIdx {
					// look upright and right
					if slices.Contains(prevLineSymbolCoords, idx+1) || slices.Contains(lineSymbolCoords, idx+1) {
						total += linePartNumber.number
					}
				}
			}
		}

		// make curr prev so next iter can make n+1 curr to compare
		prevLineParts = mapLineParts(line)
		prevLineSymbolCoords = plotLineSymbols(line)
	}
	fmt.Println("part1:", total)
}

func mapLineParts(line string) []part {
	partNumbers := []part{}
	coords := []int{}
	strChunk := ""
	for i, r := range line {
		s := string(r)
		_, err := strconv.Atoi(s)
		if err == nil {
			strChunk += s
			coords = append(coords, i)
		} else {
			partNumberNumber, err := strconv.Atoi(strChunk)
			if err == nil {
				partNumbers = append(partNumbers, part{number: partNumberNumber, coords: coords})
			}
			// reset string and idxs when hit any non-number, done with this part number
			strChunk = ""
			coords = []int{}
		}
	}
	// forgor the leftovers at end of line :skullEmoji:
	if strChunk != "" {
		partNumberNumber, err := strconv.Atoi(strChunk)
		if err == nil {
			partNumbers = append(partNumbers, part{number: partNumberNumber, coords: coords})
		}
	}
	return partNumbers
}

func plotLineSymbols(line string) []int {
	coords := []int{}
	for i, r := range line {
		s := string(r)
		if s != "." {
			_, err := strconv.Atoi(s)
			if err != nil {
				coords = append(coords, i)
			}
		}
	}
	return coords
}
