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
		fmt.Println("Total occurrences of XMAS:", countXMAS(grid))
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

func countXMAS(grid [][]rune) int {
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
				if checkWord(grid, r, c, dir[0], dir[1], "XMAS") {
					count++
				}
			}
		}
	}

	return count
}

func checkWord(grid [][]rune, r, c, dr, dc int, word string) bool {
	for i, char := range word {
		nr, nc := r+i*dr, c+i*dc
		if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) || grid[nr][nc] != char {
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
    // Check that we have enough room for the diagonals
    if r-1 < 0 || r+1 >= len(grid) || c-1 < 0 || c+1 >= len(grid[0]) {
        return false
    }

    // First diagonal: from top-left to bottom-right
    diag1 := string([]rune{grid[r-1][c-1], grid[r][c], grid[r+1][c+1]})
    // Second diagonal: from top-right to bottom-left
    diag2 := string([]rune{grid[r-1][c+1], grid[r][c], grid[r+1][c-1]})

    diag1_match := diag1 == "MAS" || diag1 == "SAM"
    diag2_match := diag2 == "MAS" || diag2 == "SAM"

    return diag1_match && diag2_match
}


func checkDiagonal(grid [][]rune, r, c, dr1, dc1, dr2, dc2 int) bool {
	// Top-Left to Bottom-Right diagonal
	if !inBounds(grid, r+dr1, c+dc1) || !inBounds(grid, r+2*dr1, c+2*dc1) {
		return false
	}
	topLeft := string([]rune{grid[r+dr1][c+dc1], grid[r][c], grid[r+dr2][c+dc2]})

	// Top-Right to Bottom-Left diagonal
	if !inBounds(grid, r+dr2, c+dc2) || !inBounds(grid, r+2*dr2, c+2*dc2) {
		return false
	}
	topRight := string([]rune{grid[r+dr2][c+dc2], grid[r][c], grid[r+dr1][c+dc1]})

	return (topLeft == "MAS" || topLeft == "SAM") && (topRight == "MAS" || topRight == "SAM")
}

func inBounds(grid [][]rune, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}
