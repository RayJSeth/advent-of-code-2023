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
	lines, err := util.ReadFileToLines("./day3/input", true)
	if err != nil {
		log.Fatal("Input file not found")
	}
	fmt.Println("Day3:")
	part1(lines)
	part2(lines)

	fmt.Print("\n")
}

func part1(lines []string) {
	total := 0

	for i := 0; i < len(lines); i++ {
		currentLine := lines[i]
		parts := mapLineParts(currentLine)
		currentSymbols := plotLineSymbols(currentLine)
		var prevSymbols []int
		var nextSymbols []int

		if i > 0 {
			prevSymbols = plotLineSymbols(lines[i-1])
		}
		if i < len(lines)-1 {
			nextSymbols = plotLineSymbols(lines[i+1])
		}

		for _, part := range parts {
			for _, partCoord := range part.coords {
				isStartOfLine := partCoord == 0
				isEndOfLine := partCoord == len(currentLine)-1
				res := scanAdjacentSymbols(part, currentSymbols, prevSymbols, nextSymbols, isStartOfLine, isEndOfLine)
				// if hit, then move to next part number
				if res > 0 {
					total += res
					break
				}
			}
		}
	}
	fmt.Println("part1:", total)
}

func part2(lines []string) {
	total := 0

	for i := 0; i < len(lines); i++ {
		currentLine := lines[i]
		gearCoords := plotLineGears(currentLine)
		// only scan when on a line with gears
		if len(gearCoords) > 0 {
			currentParts := mapLineParts(currentLine)
			var prevParts []part
			var nextParts []part

			if i > 0 {
				prevParts = mapLineParts(lines[i-1])
			}
			if i < len(lines)-1 {
				nextParts = mapLineParts(lines[i+1])
			}

			for _, gearCoord := range gearCoords {
				isStartOfLine := gearCoord == 0
				isEndOfLine := gearCoord == len(currentLine)-1

				total += scanAdjacentParts(gearCoord, currentParts, prevParts, nextParts, isStartOfLine, isEndOfLine)
			}
		}
	}

	fmt.Println("part2:", total)
}

func scanAdjacentSymbols(currentPart part, currentSymbols []int, prevSymbols []int, nextSymbols []int, isStartOfLine bool, isEndOfLine bool) int {
	for _, idx := range currentPart.coords {
		isLeftmostIdx := slices.Index(currentPart.coords, idx) == 0
		isRightmostIdx := currentPart.coords[len(currentPart.coords)-1] == idx

		// only look left if there is a column left and this is the leftmost rune of a part sequence
		if !isStartOfLine && isLeftmostIdx {
			if slices.Contains(currentSymbols, idx-1) {
				return currentPart.number
			}
		}
		// only look right if there is a column right and this is the rightmost rune of a part sequence
		if !isEndOfLine && isRightmostIdx {
			if slices.Contains(currentSymbols, idx+1) {
				return currentPart.number
			}
		}

		// only look down if there is a row below
		if nextSymbols != nil {
			// look down
			if slices.Contains(nextSymbols, idx) {
				return currentPart.number
			}
			// only look downleft if there is a column left and this is the leftmost rune of a part sequence
			if !isStartOfLine && isLeftmostIdx {
				// look downleft
				if slices.Contains(nextSymbols, idx-1) {
					return currentPart.number
				}
			}
			// only look downright if there is a column right and this is the rightmost rune of a part sequence
			if !isEndOfLine && isRightmostIdx {
				// look downright
				if slices.Contains(nextSymbols, idx+1) {
					return currentPart.number
				}
			}
		}
		// only look up if there is a row above
		if prevSymbols != nil {
			// look up
			if slices.Contains(prevSymbols, idx) {
				return currentPart.number
			}

			// only look upleft if there is a column left and this is the leftmost rune of a part sequence
			if !isStartOfLine && isLeftmostIdx {
				// look upleft and left
				if slices.Contains(prevSymbols, idx-1) {
					return currentPart.number
				}
			}
			// only look upright if there is a column right and this is the rightmost rune of a part sequence
			if !isEndOfLine && isRightmostIdx {
				// look upright and right
				if slices.Contains(prevSymbols, idx+1) {
					return currentPart.number
				}
			}
		}
	}

	return 0
}

func scanAdjacentParts(gearCoord int, currentParts []part, prevParts []part, nextParts []part, isStartOfLine bool, isEndOfLine bool) int {
	multiplicand := 0

	for _, part := range currentParts {
		// only look left if there is a column left
		if !isStartOfLine {
			if slices.Contains(part.coords, gearCoord-1) {
				if multiplicand == 0 {
					multiplicand = part.number
				} else {
					return multiplicand * part.number
				}
			}
		}
		// only look right if there is a column right
		if !isEndOfLine {
			if slices.Contains(part.coords, gearCoord+1) {
				if multiplicand == 0 {
					multiplicand = part.number
				} else {
					return multiplicand * part.number
				}
			}
		}
	}
	// only look up if there is a row up
	for _, part := range prevParts {
		// look straight up
		if slices.Contains(part.coords, gearCoord) {
			if multiplicand == 0 {
				multiplicand = part.number
				// move to next part on first hit
				// avoids double counting when gear adjacent to two numbers in same part
				continue
			} else {
				return multiplicand * part.number
			}
		}
		// only look upleft if there is a column left
		if !isStartOfLine {
			if slices.Contains(part.coords, gearCoord-1) {
				if multiplicand == 0 {
					multiplicand = part.number
					// move to next part on first hit
					// avoids double counting when gear adjacent to two numbers in same part
					continue
				} else {
					return multiplicand * part.number
				}
			}
		}
		// only look upright if there is a column right
		if !isEndOfLine {
			if slices.Contains(part.coords, gearCoord+1) {
				if multiplicand == 0 {
					multiplicand = part.number
				} else {
					return multiplicand * part.number
				}
			}
		}
	}
	// only look down if there is a row down
	for _, part := range nextParts {
		// look straight down
		if slices.Contains(part.coords, gearCoord) {
			if multiplicand == 0 {
				multiplicand = part.number
				// move to next part on first hit
				// avoids double counting when gear adjacent to two numbers in same part
				continue
			} else {
				return multiplicand * part.number
			}
		}
		// only look downleft if there is a column left
		if !isStartOfLine {
			if slices.Contains(part.coords, gearCoord-1) {
				if multiplicand == 0 {
					multiplicand = part.number
					// move to next part on first hit
					// avoids double counting when gear adjacent to two numbers in same part
					continue
				} else {
					return multiplicand * part.number
				}
			}
		}
		// only look downright if there is a column right
		if !isEndOfLine {
			if slices.Contains(part.coords, gearCoord+1) {
				if multiplicand == 0 {
					multiplicand = part.number
				} else {
					return multiplicand * part.number
				}
			}
		}
	}

	return 0
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

func plotLineGears(line string) []int {
	coords := []int{}
	for i, r := range line {
		s := string(r)
		if s == "*" {
			coords = append(coords, i)
		}
	}
	return coords
}
