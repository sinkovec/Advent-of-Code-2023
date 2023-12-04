package main

import (
	"fmt"
	"regexp"
	"sinkovec/Advent-of-Code-2023/util"
	"strconv"
)

func main() {
	data := util.ReadStringPerLine("day03/input.txt")
	result := task01(data)
	fmt.Println(result)

	result = task02(data)
	fmt.Println(result)
}

var digitRegex, _ = regexp.Compile(`\d`)
var numRegex, _ = regexp.Compile(`\d+`)
var symbolRegex, _ = regexp.Compile(`[^0-9.]`)

func task01(data []string) int {
	result := 0
	for i, line := range data {
		numLocs := numRegex.FindAllStringIndex(line, -1)
		symbolLine := digitRegex.ReplaceAllLiteralString(line, ".")
		if i > 0 {
			symbolLine = merge(symbolLine, digitRegex.ReplaceAllLiteralString(data[i-1], "."))
		}
		if i < len(data)-1 {
			symbolLine = merge(symbolLine, digitRegex.ReplaceAllLiteralString(data[i+1], "."))
		}
		symbolLocs := symbolRegex.FindAllStringIndex(symbolLine, -1)

	numLoop:
		for _, numLoc := range numLocs {
			numStart := numLoc[0]
			numEnd := numLoc[1]
			for _, symbolLoc := range symbolLocs {
				symbolPos := symbolLoc[0]

				if numStart-1 <= symbolPos && symbolPos <= numEnd {
					num, _ := strconv.Atoi(line[numStart:numEnd])
					result += num
					continue numLoop
				}
			}
		}
	}
	return result
}

func merge(line1 string, line2 string) string {
	result := ""
	for i := 0; i < len(line1); i++ {
		if symbolRegex.MatchString(string(line1[i])) || symbolRegex.MatchString(string(line2[i])) {
			result += "*"
		} else {
			result += "."
		}
	}
	return result
}

func task02(data []string) int {
	result := 0
	return result
}
