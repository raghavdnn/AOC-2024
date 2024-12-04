package main

import (
	"fmt"
	"strings"
	"aoc/utils"
)

func main() {
	part := utils.ParseFlags()

	lines, err := utils.ReadFileLines("day2/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	safeReports := 0
	for _, line := range lines {
		parts := strings.Fields(line)
		levels, err := utils.ParseIntSlice(parts)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch part {
		case "a":
			if isSafeReport(levels) {
				safeReports++
			}
		case "b":
			if isSafeReport(levels) || isSafeWithDampener(levels) {
				safeReports++
			}
		}
	}

	fmt.Println("Number of safe reports:", safeReports)
}

func isSafeReport(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	isIncreasing := levels[1] > levels[0]
	isDecreasing := levels[1] < levels[0]

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff < -3 || diff > 3 {
			return false
		}
		if isIncreasing && levels[i] <= levels[i-1] {
			return false
		}
		if isDecreasing && levels[i] >= levels[i-1] {
			return false
		}
		if levels[i] == levels[i-1] {
			return false
		}
	}

	return true
}

func isSafeWithDampener(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		modifiedLevels := append([]int{}, levels[:i]...)
		modifiedLevels = append(modifiedLevels, levels[i+1:]...)
		if isSafeReport(modifiedLevels) {
			return true
		}
	}
	return false
}
