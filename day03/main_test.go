package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3_Task01(t *testing.T) {
	result := task01([]string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	})
	assert.Equal(t, 4361, result)
}

func TestDay3_Task02(t *testing.T) {
	result := task02([]string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	})
	assert.Equal(t, 0, result)
}
