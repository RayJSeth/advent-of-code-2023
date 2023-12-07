package util

import (
	"os"
	"strings"
)

func ReadFileToLines(path string) ([]string, error) {
	in, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(in), "\n")
	return lines[:len(lines)-1], nil
}
