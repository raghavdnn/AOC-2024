package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	part := utils.ParseFlags()

	lines, err := utils.ReadFileLines("day4/input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	grid := parseGrid(lines)

	switch part {
	case "a":
		fmt.Println("Total occurrences of XMAS:", countWord(grid, "XMAS"))
	case "b":
		fmt.Println("Total occurrences of X-MAS:", countXMasPattern(grid))
	}
}

func parseGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func countWord(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	directions := [][2]int{
		{0, 1},   // Right
		{0, -1},  // Left
		{1, 0},   // Down
		{-1, 0},  // Up
		{1, 1},   // Down-Right
		{-1, -1}, // Up-Left
		{1, -1},  // Down-Left
		{-1, 1},  // Up-Right
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				dr, dc := dir[0], dir[1]
				// Pre-check bounds based on word length, skip redundant iterations boi
				if inBounds(grid, r, c) && inBounds(grid, r+(len(word)-1)*dr, c+(len(word)-1)*dc) {
					if checkWord(grid, r, c, dr, dc, word) {
						count++
					}
				}
			}
		}
	}

	return count
}

func checkWord(grid [][]rune, r, c, dr, dc int, word string) bool {
	for i, char := range word {
		nr, nc := r+i*dr, c+i*dc
		if !inBounds(grid, nr, nc) || grid[nr][nc] != char {
			return false
		}
	}
	return true
}

func countXMasPattern(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if isXMasPattern(grid, r, c) {
				count++
			}
		}
	}

	return count
}

func isXMasPattern(grid [][]rune, r, c int) bool {
	if !inBounds(grid, r-1, c-1) || !inBounds(grid, r+1, c+1) ||
		!inBounds(grid, r-1, c+1) || !inBounds(grid, r+1, c-1) {
		return false
	}

	diag1 := string([]rune{grid[r-1][c-1], grid[r][c], grid[r+1][c+1]})
	diag2 := string([]rune{grid[r-1][c+1], grid[r][c], grid[r+1][c-1]})

	diag1_match := diag1 == "MAS" || diag1 == "SAM"
	diag2_match := diag2 == "MAS" || diag2 == "SAM"

	return diag1_match && diag2_match
}

func inBounds(grid [][]rune, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}
