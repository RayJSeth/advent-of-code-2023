package day5

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"rayjseth.io/advent2023/util"
)

func GetAnswers() {
	lines, err := util.ReadFileToLines("./day5/input", false)
	if err != nil {
		log.Fatal("Input file not found")
	}
	fmt.Println("Day5:")
	part1(lines)
	fmt.Print("\n")
}

type mapEntry struct {
	dStart int
	sStart int
	mRange int
}

func part1(lines []string) {
	seeds := strings.Split(strings.Split(lines[0], ": ")[1], " ")

	// starting coords are just the seeds
	var coords []int
	for _, seed := range seeds {
		sNum, _ := strconv.Atoi(seed)
		coords = append(coords, sNum)
	}

	me := []mapEntry{}

	for i := 3; i < len(lines); i++ {
		line := lines[i]
		if line != "" {
			s := string(line[0])
			_, err := strconv.Atoi(s)
			// we're still in the map of numbers so keep parsing
			if err == nil {
				meArr := strings.Split(line, " ")
				dStart, _ := strconv.Atoi(meArr[0])
				sStart, _ := strconv.Atoi(meArr[1])
				mRange, _ := strconv.Atoi(meArr[2])
				me = append(me, mapEntry{dStart: dStart, sStart: sStart, mRange: mRange})
			} else {
				// reset the parsed map since this is the label for the next map
				me = nil
			}
			// empty line means current map is parsed so now interpret results
		} else {
			var nextCoords []int
			for _, point := range coords {
				// by default points map to themself if no hit
				hit := point
				for _, e := range me {
					// if the point is within the sourceStart + sourceEnd (sStart + mRange) it's a hit
					if point >= e.sStart && point <= e.sStart+e.mRange {
						offset := e.dStart - e.sStart
						hit = point + offset
					}
				}
				nextCoords = append(nextCoords, hit)
			}
			// advance the coordinates
			coords = nextCoords
		}
	}

	// find the lowest final coordinate
	lowestLoc := coords[0]
	for i := 1; i < len(coords); i++ {
		coord := coords[i]
		if coord < lowestLoc {
			lowestLoc = coord
		}
	}

	fmt.Println("Part1:", lowestLoc)
}
