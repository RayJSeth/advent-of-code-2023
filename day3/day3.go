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

		// anchor on previous line to look for symbols downwards or laterally
		// a little awkward with the asymmetry - only one of these up or down needs to do the lateral
		// just doing it here since it comes first
		for _, prevLinePartNumber := range prevLineParts {
			for _, idx := range prevLinePartNumber.coords {
				isLeftmostIdx := slices.Index(prevLinePartNumber.coords, idx) == 0
				isRightmostIdx := prevLinePartNumber.coords[len(prevLinePartNumber.coords)-1] == idx
				isOnLastLine := !(i < len(lines))

				// always look down on nonlast line
				if !isOnLastLine {
					if slices.Contains(lineSymbolCoords, idx) {
						total += prevLinePartNumber.number
					}
				}

				// is margin left and cursor left
				if idx > 0 && isLeftmostIdx {
					// look left
					if slices.Contains(prevLineSymbolCoords, idx-1) {
						total += prevLinePartNumber.number
						// is margin down
					} else if !isOnLastLine {
						// look downleft
						if slices.Contains(lineSymbolCoords, idx-1) {
							total += prevLinePartNumber.number
						}
					}
				}
				// is margin right and cursor right
				if idx < len(line) && isRightmostIdx {
					// look downright and right
					if slices.Contains(lineSymbolCoords, idx+1) || slices.Contains(prevLineSymbolCoords, idx+1) {
						total += prevLinePartNumber.number
					}
				}

			}
		}

		// anchor on current line to look for symbols upwards
		for _, linePartNumber := range lineParts {
			for _, idx := range linePartNumber.coords {
				isLeftmostIdx := slices.Index(linePartNumber.coords, idx) == 0
				isRightmostIdx := linePartNumber.coords[len(linePartNumber.coords)-1] == idx

				// always look up since start on 1th line
				if slices.Contains(prevLineSymbolCoords, idx) {
					total += linePartNumber.number
				}

				// is margin left and cursor left
				if idx > 0 && isLeftmostIdx {
					// look upleft
					if slices.Contains(prevLineSymbolCoords, idx-1) {
						total += linePartNumber.number
					}
				}
				// is margin right and cursor right
				if idx < len(line) && isRightmostIdx {
					// look upright
					if slices.Contains(prevLineSymbolCoords, idx+1) {
						total += linePartNumber.number
					}
				}
			}
		}

		// make curr prev so next iter can make n+1 curr to compare
		prevLine = line
		prevLineParts = mapLineParts(prevLine)
		prevLineSymbolCoords = plotLineSymbols(prevLine)
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
