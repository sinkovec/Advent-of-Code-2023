package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1_Task01(t *testing.T) {
	result := task02([]string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet"})
	assert.Equal(t, 142, result)
}

func TestDay1_Task02(t *testing.T) {
	result := task02([]string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen"})
	assert.Equal(t, 281, result)
}
