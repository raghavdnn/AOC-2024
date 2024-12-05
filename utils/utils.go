package utils

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseFlags() string {
	part := flag.String("part", "", "Specify the part to solve: a or b")
	flag.Parse()

	if *part != "a" && *part != "b" {
		fmt.Println("Error: You must specify the part to solve using -part=a or -part=b")
		os.Exit(1)
	}

	return *part
}

func ReadFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening input file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input file: %v", err)
	}

	return lines, nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ParseIntSlice(parts []string) ([]int, error) {
	ints := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("error parsing number: %v", err)
		}
		ints[i] = num
	}
	return ints, nil
}

func JoinLines(lines []string) string {
	return strings.Join(lines, "")
}

func SplitByEmptyLines(lines []string) [][]string {
    var sections [][]string
    start := 0

    for i, line := range lines {
        if line == "" {
            if start < i {
                sections = append(sections, lines[start:i])
            }
            start = i + 1
        }
    }

    if start < len(lines) {
        sections = append(sections, lines[start:])
    }

    return sections
}

