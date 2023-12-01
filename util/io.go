package util

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("could not open %s: %v", path, err)
	}
	return string(data)
}

func ReadStringPerLine(path string) []string {
	return ReadStringWithDelimiter(path, "\n")
}

func ReadStringWithDelimiter(path string, delimiter string) []string {
	data := ReadInput(path)
	result := make([]string, 0)
	for _, s := range strings.Split(data, delimiter) {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		result = append(result, s)
	}
	return result
}

func ReadIntPerLine(path string) []int {
	return ReadIntWithDelimiter(path, "\n")
}

func ReadInts(path string) []int {
	return ReadIntWithDelimiter(path, ",")
}

func ReadIntWithDelimiter(path string, delimiter string) []int {
	data := ReadStringWithDelimiter(path, delimiter)
	result := make([]int, len(data))
	for i, s := range data {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil // fmt.Errorf("failed to convert string '%s' to int: %v", s, err)
		}
		result[i] = num
	}
	return result
}
