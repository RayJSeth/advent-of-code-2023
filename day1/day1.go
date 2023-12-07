package day1

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"rayjseth.io/advent2023/util"
)

var conversions = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func GetAnswers() {
	lines, err := util.ReadFileToLines("./day1/input")
	if err != nil {
		log.Fatal("Input file not found")
	}
	fmt.Println("Day1:")
	part1(lines)
	part2(lines)
	fmt.Print("\n")
}

func part1(lines []string) {
	total := 0

	// Overkill? Maybe. Fun? Yes.
	var wg sync.WaitGroup
	resultCh := make(chan int, len(lines))

	for _, line := range lines {
		wg.Add(1)
		go processLine(line, false, &wg, resultCh)
	}

	wg.Wait()
	close(resultCh)

	for result := range resultCh {
		total += result
	}

	fmt.Println("part1:", total)
}

func part2(lines []string) {
	total := 0

	// Overkill? Maybe. Fun? Yes.
	var wg sync.WaitGroup
	resultCh := make(chan int, len(lines))

	for _, line := range lines {
		wg.Add(1)
		go processLine(line, true, &wg, resultCh)
	}

	wg.Wait()
	close(resultCh)

	for result := range resultCh {
		total += result
	}

	fmt.Println("part2:", total)
}

func processLine(line string, convertTextNum bool, wg *sync.WaitGroup, resultCh chan<- int) {
	defer wg.Done()

	resultCh <- 10*getFirstCoord(line, convertTextNum) + getLastCoord(line, convertTextNum)
}

func getFirstCoord(s string, convertTextNum bool) int {
	firstCoord := 0
	strChunk := ""

strLoop:
	for i := 0; i < len(s); i++ {
		cs := string(s[i])

		num, err := strconv.Atoi(cs)
		if err == nil {
			firstCoord = num
			break strLoop
		}

		if convertTextNum {
			strChunk += cs

			for k, v := range conversions {
				if strings.Contains(strChunk, k) {
					firstCoord = v
					break strLoop
				}
			}
		}
	}

	return firstCoord
}

func getLastCoord(s string, convertTextNum bool) int {
	lastCoord := 0
	strChunk := ""

strLoop:
	for i := len(s) - 1; i >= 0; i-- {
		cs := string(s[i])

		num, err := strconv.Atoi(cs)
		if err == nil {
			lastCoord = num
			break strLoop
		}

		if convertTextNum {
			// gotta prepend this since walking backwards
			strChunk = cs + strChunk

			for k, v := range conversions {
				if strings.Contains(strChunk, k) {
					lastCoord = v
					break strLoop
				}
			}
		}
	}

	return lastCoord
}
