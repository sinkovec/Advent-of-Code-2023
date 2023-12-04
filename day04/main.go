package main

import (
	"fmt"
	"sinkovec/Advent-of-Code-2023/util"
	"strconv"
	"strings"
)

func main() {
	data := util.ReadStringPerLine("day04/input.txt")
	result := task01(data)
	fmt.Println(result)

	result = task02(data)
	fmt.Println(result)
}

func task01(data []string) int {
	result := 0
	for i, card := range data {
		cardId := i + 1
		card, _ = strings.CutPrefix(card, "Card "+strconv.Itoa(cardId)+": ")
		winningNumbers := map[string]struct{}{}
		draw := strings.Split(card, " | ")
		matches := 0
		for _, winningNumber := range strings.Split(draw[0], " ") {
			winningNumber = strings.TrimSpace(winningNumber)
			if winningNumber == "" {
				continue
			}
			winningNumbers[winningNumber] = struct{}{}
		}
		for _, number := range strings.Split(draw[1], " ") {
			number = strings.TrimSpace(number)
			if _, ok := winningNumbers[number]; ok {
				if matches == 0 {
					matches = 1
				} else {
					matches *= 2
				}
			}
		}
		result += matches
	}
	return result
}

type stack []int

func (s stack) push(v int) stack {
	return append(s, v)
}

func (s stack) pop() (stack, int) {
	return s[1:], s[0]
}

func task02(data []string) int {
	result := 0
	copies := make(stack, 0)
	for i, card := range data {
		cardId := i + 1
		card, _ = strings.CutPrefix(card, "Card "+strconv.Itoa(cardId)+": ")
		winningNumbers := map[string]struct{}{}
		draw := strings.Split(card, " | ")
		matches := 0
		for _, winningNumber := range strings.Split(draw[0], " ") {
			winningNumber = strings.TrimSpace(winningNumber)
			if winningNumber == "" {
				continue
			}
			winningNumbers[winningNumber] = struct{}{}
		}
		for _, number := range strings.Split(draw[1], " ") {
			number = strings.TrimSpace(number)
			if _, ok := winningNumbers[number]; ok {
				matches++
			}
		}
		cardInstances := 1
		if len(copies) > 0 {
			newCopies, copiesOfCard := copies.pop()
			copies = newCopies
			cardInstances += copiesOfCard
		}
		copySize := len(copies)
		for i := 0; i < matches; i++ {
			if i < copySize {
				copies[i] += cardInstances
			} else {
				copies = copies.push(cardInstances)
			}
		}
		result += cardInstances
	}
	return result
}
