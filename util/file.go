package util

import (
	"os"
	"strings"
)

func ReadFileToLines(path string, trimLastNewline bool) ([]string, error) {
	in, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(in), "\n")
	if trimLastNewline {
		lines = lines[:len(lines)-1]
	}
	return lines, nil
}
