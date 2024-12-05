package puzzles

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Day 5: Print Queue
// https://adventofcode.com/2024/day/5
func dayFive(input string) (int, int) {
	return d5p1(input), 0
}

// Completes the first half of the puzzle for day 4.
func d5p1(input string) int {
	checksum := 0
	validUpdates := make([][]int, 0)
	dependencies, updates := parseManual(input)

	for _, update := range updates {
		valid := true
		for i, page := range update {
			dependants, exists := dependencies[page]

			// This page has no dependants, no need to validate it.
			if !exists {
				continue
			}

			for _, dependant := range dependants {
				for j := 0; j < i+1; j++ {
					if update[j] == dependant {
						valid = false
						goto dealWithUpdate
					}
				}
			}
		}

	dealWithUpdate:
		if valid {
			validUpdates = append(validUpdates, update)
		}
	}

	for _, update := range validUpdates {
		middle := (len(update) - 1) / 2
		checksum += update[middle]
	}

	return checksum
}

// Parses the input file into usable data.
//
// Returns a tuple, where:
// - the first parameter is a map, keyed by page values, containing a slice of dependant pages;
// - the second parameter is a slice of updates, split into page numbers.
func parseManual(input string) (map[int][]int, [][]int) {
	var updates [][]int
	dependencies := make(map[int][]int)

	reDeps := regexp.MustCompile(`\d+\|\d+`)
	deps := reDeps.FindAllString(input, -1)
	for _, dep := range deps {
		instruction := strings.Split(dep, "|")
		page, _ := strconv.Atoi(instruction[0])
		dependant, _ := strconv.Atoi(instruction[1])
		_, exists := dependencies[page]
		if exists {
			dependencies[page] = append(dependencies[page], dependant)
		} else {
			dependencies[page] = []int{dependant}
		}
	}

	// Sort the dependants to make lookups faster.
	for page, _ := range dependencies {
		slices.Sort(dependencies[page])
	}

	reUps := regexp.MustCompile(`(\d+,)+\d+`)
	ups := reUps.FindAllString(input, -1)
	updates = make([][]int, len(ups))
	for i, up := range ups {
		pages := strings.Split(up, ",")
		updates[i] = make([]int, len(pages))
		for j, page := range pages {
			updates[i][j], _ = strconv.Atoi(page)
		}
	}

	return dependencies, updates
}
