package main

import (
    "aoc/utils"
    "fmt"
    "sort"
    "strconv"
    "strings"
)

func main() {
    part := utils.ParseFlags()

    lines, _ := utils.ReadFileLines("day5/input.txt")

    sections := utils.SplitByEmptyLines(lines)

    if len(sections) < 2 {
        fmt.Println("Input file should contain at least two sections.")
        return
    }

    rules := parseOrderingRules(sections[0])

    fmt.Println("Parsed Ordering Rules:")
    for key, value := range rules {
        fmt.Println("Page", key, "must come before pages", value)
    }
	fmt.Println()

	switch part {
		case "a":
			fmt.Println("\nSum of middle pages from valid updates:", processUpdates(sections[1], rules, false))
		case "b":
			fmt.Println("\nSum of middle pages from reordered updates:", processUpdates(sections[1], rules, true))
	}
}

func parseOrderingRules(lines []string) map[int][]int {
    rules := make(map[int][]int)
    for _, line := range lines {
        parts := strings.Split(line, "|")
        if len(parts) != 2 {
            continue
        }
        x, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
        y, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
        if err1 != nil || err2 != nil {
            continue
        }
        rules[x] = append(rules[x], y)
    }
    return rules
}

func processUpdates(updates []string, rules map[int][]int, reorder bool) int {
    total := 0
    for i, line := range updates {
        fmt.Printf("Processing Update %d: %s\n", i+1, line)
        update, err := parseUpdate(line)
        if err != nil {
            fmt.Println("Error parsing update:", err)
            fmt.Println()
            continue
        }

        isCorrectOrder := isUpdateInCorrectOrder(update, rules)

        if !isCorrectOrder {
            fmt.Println("Update is NOT in correct order.")
            if reorder {
                sortedUpdate, err := topologicalSort(update, rules)
                if err != nil {
                    fmt.Println("Error sorting update:", err)
                    fmt.Println()
                    continue
                }
                fmt.Printf("Reordered Update: %v\n", sortedUpdate)
                update = sortedUpdate
            } else {
                fmt.Println()
                continue
            }
        } else {
            fmt.Println("Update is in correct order.")
            if reorder {
                fmt.Println()
                continue
            }
        }

        middlePage := getMiddlePage(update)
        fmt.Printf("Middle page of the update: %d\n", middlePage)
        total += middlePage
        fmt.Println()
    }
    return total
}


func parseUpdate(line string) ([]int, error) {
    parts := strings.Split(line, ",")
    update := make([]int, 0, len(parts))
    pageSet := make(map[int]bool)
    for _, part := range parts {
        num, err := strconv.Atoi(strings.TrimSpace(part))
        if err != nil {
            return nil, fmt.Errorf("invalid page number '%s'", part)
        }
        if pageSet[num] {
            return nil, fmt.Errorf("duplicate page number '%d' in update", num)
        }
        pageSet[num] = true
        update = append(update, num)
    }
    return update, nil
}

func isUpdateInCorrectOrder(update []int, rules map[int][]int) bool {
    position := make(map[int]int)
    for idx, page := range update {
        position[page] = idx
    }

    for x := range position {
        ys, xHasRules := rules[x]
        if !xHasRules {
            continue
        }
        xPos := position[x]
        for _, y := range ys {
            yPos, yExists := position[y]
            if !yExists {
                continue
            }
            if xPos >= yPos {
                fmt.Printf("Ordering violation: Page %d should come before page %d\n", x, y)
                return false
            }
        }
    }
    return true
}

// Kahn you believe it? We're sorting topologically!
// Borrowed some wisdom for this topological sort!
func topologicalSort(pages []int, rules map[int][]int) ([]int, error) {
    // Build the subgraph
    graph := make(map[int][]int)
    inDegree := make(map[int]int)

    pageSet := make(map[int]bool)
    for _, page := range pages {
        pageSet[page] = true
        graph[page] = []int{}
        inDegree[page] = 0
    }

    // Build the graph based on ordering rules
    for x, ys := range rules {
        if !pageSet[x] {
            continue
        }
        for _, y := range ys {
            if !pageSet[y] {
                continue
            }
            graph[x] = append(graph[x], y)
            inDegree[y]++
        }
    }

    // Initialize the zero in-degree list
    zeroInDegree := make([]int, 0)
    for node, deg := range inDegree {
        if deg == 0 {
            zeroInDegree = append(zeroInDegree, node)
        }
    }
    // Sort in reverse order to prioritize higher page numbers
    sort.Sort(sort.Reverse(sort.IntSlice(zeroInDegree)))

    // Kahn's Algorithm for topological sorting
    var sorted []int
    for len(zeroInDegree) > 0 {
        node := zeroInDegree[0]
        zeroInDegree = zeroInDegree[1:]
        sorted = append(sorted, node)

        for _, neighbor := range graph[node] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                zeroInDegree = append(zeroInDegree, neighbor)
                // Keep zeroInDegree sorted in reverse order
                sort.Sort(sort.Reverse(sort.IntSlice(zeroInDegree)))
            }
        }
    }

    if len(sorted) != len(pages) {
        return nil, fmt.Errorf("cycle detected, cannot perform topological sort")
    }

    return sorted, nil
}

func getMiddlePage(update []int) int {
    n := len(update)
    if n%2 == 0 {
        return update[n/2-1]
    }
    return update[n/2]
}
