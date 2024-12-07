package puzzles

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

type pointingDirection int

// Enum for directions the guard could face.
const (
	Up pointingDirection = iota
	Down
	Left
	Right
)

var pointing = map[pointingDirection]string{
	Up:    "up",
	Down:  "down",
	Left:  "left",
	Right: "right",
}

// Handles rotation based on the current facing direction.
func rotate(p pointingDirection) pointingDirection {
	switch p {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}

	panic("Invalid pointing direction!")
}

// Determines the y and x moving speed based on the direction faced.
func moving(p pointingDirection) (int, int) {
	var (
		yd, xd int
	)

	switch p {
	case Up:
		yd = -1
		xd = 0
	case Down:
		yd = 1
		xd = 0
	case Left:
		yd = 0
		xd = -1
	case Right:
		yd = 0
		xd = 1
	}

	return yd, xd
}

// Day 6: Guard Gallivant
// https://adventofcode.com/2024/day/6
func daySix(input string) (int, int) {
	return d6p1(input), d6p2(input)
}

// Completes the first half of the puzzle for day 6.
func d6p1(input string) int {
	room, guard := parseLaboratory(input)
	visited, _, _ := walkPath(room, guard, Up)
	totalVisits := 0

	for _, tiles := range visited {
		totalVisits += len(tiles)
	}

	return totalVisits
}

// Completes the second half of the puzzle for day 6.
func d6p2(input string) int {
	room, guard := parseLaboratory(input)
	wouldVisit, _, _ := walkPath(room, guard, Up)
	loops := 0

	for y, tiles := range wouldVisit {
		for _, x := range tiles {
			// Don't put an object where the guard is supposed to start.
			if y == guard[0] && x == guard[1] {
				continue
			}

			// Place the object in front of the path tile.
			altRoom := make([][]byte, len(room))
			for i, row := range room {
				altRoom[i] = make([]byte, len(row))
				copy(altRoom[i], row)
			}
			altRoom[y][x] = '#'

			// Determine if a loop was created.
			_, _, err := walkPath(altRoom, guard, Up, true)
			if err != nil {
				loops++
			}
		}
	}

	return loops
}

// Walks the path around the provided room from the provided starting location.
// Returns:
//   - A map of visited coordinates, where the key is the y coordinate
//     and the value is a slice of corresponding x coordinates;
//   - A map of faced directions for a coordinate, where the key is the y coordinate,
//     the value is another map, where the key is the corresponding x coordinate,
//     and the value of the second map is a slice of directions faced on this tile;
//   - An error, in case a loop was detected.
func walkPath(
	room [][]byte,
	guard [2]int,
	pointing pointingDirection,
	debug ...bool,
) (
	map[int][]int,
	map[int]map[int][]pointingDirection,
	error,
) {
	visited := make(map[int][]int)
	faced := make(map[int]map[int][]pointingDirection)

	for true {
		yd, xd := moving(pointing)

		for steps := 0; true; steps++ {
			x := guard[1] + xd*steps
			y := guard[0] + yd*steps
			if debug != nil {
				room[y][x] = 'X'
			}

			// If this is a new tile, add it to the mapping.
			_, exists := visited[y]
			if exists {
				if !slices.Contains(visited[y], x) {
					// This is the first time we've visited this tile.
					visited[y] = append(visited[y], x)
					faced[y][x] = []pointingDirection{pointing}
				} else if !slices.Contains(faced[y][x], pointing) {
					// We have been here before, but haven't went this way yet.
					faced[y][x] = append(faced[y][x], pointing)
				} else {
					f, _ := os.OpenFile("./garb.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					defer f.Close()
					for _, row := range room {
						f.Write(row)
						f.WriteString("\n")
					}
					f.WriteString("\n\n")
					// We have been here while facing the same direction. We're in a loop!
					return nil, nil, fmt.Errorf("Infinite loop detected!")
				}
			} else {
				// This is the first time we've visited this row, so it's definitely a new tile.
				visited[y] = []int{x}
				faced[y] = make(map[int][]pointingDirection)
				faced[y][x] = []pointingDirection{pointing}
			}

			// Check if we're about to go out of the room.
			if y+yd < 0 || y+yd >= len(room) || x+xd < 0 || x+xd >= len(room[1]) {
				return visited, faced, nil
			}

			// If we run into a dead-end, rotate and restart our steps counter.
			if room[y+yd][x+xd] == '#' {
				pointing = rotate(pointing)
				guard[0] = y
				guard[1] = x
				break
			}
		}
	}

	return visited, faced, nil
}

// Parses the input data, returning:
// - A 2D slice of the room's layout in y,x space, replacing the guard with a period;
// - The guard's position within the room in y,x space.
func parseLaboratory(input string) ([][]byte, [2]int) {
	// Find the position of the guard within the data and remove him from the room.
	gi := strings.Index(input, "^")
	input = strings.Replace(input, "^", ".", -1)

	// Parse the room into a 2D slice of y,x coordinates.
	reRows := regexp.MustCompile(`[.#]+`)
	rows := reRows.FindAll([]byte(input), -1)
	room := make([][]byte, len(rows))
	for i, row := range rows {
		room[i] = make([]byte, len(row))
		room[i] = row
	}

	// Find the coordinates of the guard within the room.
	gy := gi / (len(rows[0]) + 1)
	gx := gi - (gy)*(len(rows[0])+1)
	guard := [2]int{gy, gx}

	return room, guard
}
