package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
)

var (
	mulRegex     = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	controlRegex = regexp.MustCompile(`(do\(\)|don't\(\))`)
)

func main() {
	part := utils.ParseFlags()

	input, err := utils.ReadFileLines("day3/input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	joinedInput := utils.JoinLines(input)

	switch part {
	case "a":
		fmt.Println("Total sum of all valid multiplications:", solvePartA(joinedInput))
	case "b":
		fmt.Println("Total sum of enabled multiplications:", solvePartB(joinedInput))
	}
}

func solvePartA(input string) int {
	matches := mulRegex.FindAllStringSubmatch(input, -1)
	total := 0
	for _, match := range matches {
		total += multiply(match[1], match[2])
	}
	return total
}

func solvePartB(input string) int {
	total := 0
	enabled := true
	pos := 0

	for pos < len(input) {
		controlMatch := controlRegex.FindStringIndex(input[pos:])
		mulMatch := mulRegex.FindStringSubmatchIndex(input[pos:])

		if controlMatch != nil && (mulMatch == nil || controlMatch[0] < mulMatch[0]) {
			enabled = input[pos+controlMatch[0]:pos+controlMatch[1]] == "do()"
			pos += controlMatch[1]
		} else if mulMatch != nil {
			if enabled {
				total += multiply(input[pos+mulMatch[2]:pos+mulMatch[3]], input[pos+mulMatch[4]:pos+mulMatch[5]])
			}
			pos += mulMatch[1]
		} else {
			break
		}
	}

	return total
}

func multiply(xStr, yStr string) int {
	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)
	return x * y
}
