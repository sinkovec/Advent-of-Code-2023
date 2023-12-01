package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	result := day01([]string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet"})
	assert.Equal(t, 142, result)
}
