package puzzles

import (
	"regexp"
	"sort"
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
	sort.Ints(left)
	sort.Ints(right)

	// Calculate the total distance between the two lists.
	distance := 0
	for i := 0; i < len(left); i++ {
		d := right[i] - left[i]
		if d < 0 {
			d = -d
		}
		distance += d
	}

	return distance, 0
}
