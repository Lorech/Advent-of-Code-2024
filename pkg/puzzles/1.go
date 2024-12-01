package puzzles

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Provides the solution for day 1.
//
// In day 1, we must help The Historians find the location IDs of historically
// significant locations where the Chief Historian could be found after disappearing.
//
// To do this, we must combine two individual lists of location IDs found by the
// historians, and and the distances between each of their IDs. We can do this
// by aligning the two lists of IDs from smallest to largest, subtracting their
// values to get individual distances between locations, and then summing the
// distances to get the total distance between the two lists.
//
// As a second form of comparison, we must find the similarity between the two
// lists - maybe some of the locations are the same, indicating that they are
// actually valid locations. We can do this by comparing the two lists of IDs,
// multiplying each ID found in the left list with the amount of times is also
// appears on the right list, repeating IDs included.
func DayOne(input string) (int, int) {
	var left, right []int

	// Parse both lists into separate integer slices.
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		ids := regexp.MustCompile(`\s+`).Split(row, -1)

		if len(ids) != 2 {
			break
		}

		first, error := strconv.Atoi(ids[0])
		if error != nil {
			panic(error)
		}
		second, error := strconv.Atoi(ids[1])
		if error != nil {
			panic(error)
		}

		left = append(left, first)
		right = append(right, second)
	}

	// Sort both lists in ascending order.
	slices.Sort(left)
	slices.Sort(right)

	// Calculate the total distance between the two lists.
	distance := 0
	for i := 0; i < len(left); i++ {
		d := right[i] - left[i]
		if d < 0 {
			d = -d
		}
		distance += d
	}

	// Calculate the similarity between the two lists.
	similarity := 0
	for _, id := range left {
		position, appears := slices.BinarySearch(right, id)
		if appears {
			c := 0
			for i := position; i < len(right); i++ {
				if right[i] != id {
					break
				}

				c++
			}
			similarity += id * c
		}
	}

	return distance, similarity
}
