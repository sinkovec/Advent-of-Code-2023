package main

import (
	"fmt"
	"sinkovec/Advent-of-Code-2023/util"
	"strconv"
	"strings"
)

func main() {
	data := util.ReadStringPerLine("day02/input.txt")
	result := task01(data)
	fmt.Println(result)

	result = task02(data)
	fmt.Println(result)
}

func task01(data []string) int {
	result := 0
	isGamePossible := true
	for i, game := range data {
		isGamePossible = true
		gameId := i + 1
		game, _ = strings.CutPrefix(game, "Game "+strconv.Itoa(gameId)+":")
		for _, draw := range strings.Split(game, ";") {
			for _, cube := range strings.Split(draw, ",") {
				numWithColor := strings.Split(strings.TrimSpace(cube), " ")
				num, _ := strconv.Atoi(numWithColor[0])
				color := numWithColor[1]
				if (color == "red" && num > 12) || (color == "green" && num > 13) || (color == "blue" && num > 14) {
					isGamePossible = false
				}
			}
		}
		if isGamePossible {
			result += gameId
		}

	}
	return result
}

func task02(data []string) int {
	result := 0
	for i, game := range data {
		gameId := i + 1
		game, _ = strings.CutPrefix(game, "Game "+strconv.Itoa(gameId)+":")
		minCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, draw := range strings.Split(game, ";") {
			for _, cube := range strings.Split(draw, ",") {
				numWithColor := strings.Split(strings.TrimSpace(cube), " ")
				num, _ := strconv.Atoi(numWithColor[0])
				color := numWithColor[1]
				minCubes[color] = max(num, minCubes[color])
			}
		}
		power := 1
		for _, num := range minCubes {
			power *= num
		}
		result += power
	}
	return result
}
