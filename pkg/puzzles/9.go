package puzzles

import (
	"fmt"
	"strconv"
	"strings"
)

// Day 9: Disk Fragmenter
// https://adventofcode.com/2024/day/9
func dayNine(input string) (int, int) {
	return 0, 0
}

// Completes the first half of the puzzle for day 9.
func d9p1(input string) int {
	disk := parseDisk(input)
	checksum := 0

	fmt.Println(disk)

	return checksum
}

// Parses the input data, converting it to an uncompressed disk drive.
func parseDisk(input string) []int {
	compression, _ := strings.CutSuffix(input, "\n")
	disk := make([]int, 0)

	for i := 0; i < len(compression); i += 2 {
		f := i / 2
		l, _ := strconv.Atoi(string(compression[i]))
		s := 0

		// Free space is only added for blocks before the final one.
		if i != len(compression)-1 {
			s, _ = strconv.Atoi(string(compression[i+1]))
		}

		// Map the file onto the disk.
		files := make([]int, l)
		for j := range files {
			files[j] = f
		}
		disk = append(disk, files...)

		// Map the free space onto the disk.
		space := make([]int, s)
		for j := range space {
			space[j] = -1
		}
		disk = append(disk, space...)
	}

	return disk
}
