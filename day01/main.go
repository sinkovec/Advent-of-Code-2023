package main

import (
	"fmt"
	"regexp"
	"sinkovec/Advent-of-Code-2023/util"
	"strconv"
)

func main() {
	data := util.ReadStringPerLine("day01/input.txt")
	result := day01(data)
	fmt.Println(result)

	result = day02(data)
	fmt.Println(result)
}

func day01(data []string) int {
	r, _ := regexp.Compile(`[0-9]`)
	result := 0
	for _, line := range data {
		ints := r.FindAllString(line, -1)
		if ints == nil {
			continue
		}
		numString := ints[0] + ints[len(ints)-1]
		i, _ := strconv.Atoi(numString)
		result += i
	}
	return result
}

func day02(data []string) int {
	return 0
}
