package main

import (
	"fmt"
	"sort"
	"strings"
	"aoc/utils"
)

func main() {
	part := utils.ParseFlags()

	lines, err := utils.ReadFileLines("day1/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var leftList, rightList []int
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid input format")
			return
		}

		nums, err := utils.ParseIntSlice(parts)
		if err != nil {
			fmt.Println(err)
			return
		}

		leftList = append(leftList, nums[0])
		rightList = append(rightList, nums[1])
	}

	switch part {
	case "a":
		fmt.Println("Total Distance:", calculateTotalDistance(leftList, rightList))
	case "b":
		fmt.Println("Similarity Score:", calculateSimilarityScore(leftList, rightList))
	}
}

func calculateTotalDistance(leftList, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += utils.Abs(leftList[i] - rightList[i])
	}
	return totalDistance
}

func calculateSimilarityScore(leftList, rightList []int) int {
	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}

	similarityScore := 0
	for _, num := range leftList {
		similarityScore += num * rightCount[num]
	}
	return similarityScore
}
