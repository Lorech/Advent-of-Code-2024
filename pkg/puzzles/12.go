package puzzles

import (
	"fmt"
	"strings"
)

type plant struct {
	kind      rune // The kind of plant growing here.
	x         int  // The x position of the plant within the garden.
	y         int  // The y position of the plant within the garden.
	neighbors int  // The number of adjacent plants to this plant.
}

// Day 12: Garden Groups
// https://adventofcode.com/2024/day/12
func dayTwelve(input string) (int, int) {
	return d12p1(input), 0
}

// Completes the first half of the puzzle for day 12.
func d12p1(input string) int {
	garden := parseGarden(input)
	plots := make([][]plant, 0)
	price := 0

	for y, row := range garden {
		for x, tile := range row {
			if tile != '.' {
				start := [2]int{y, x}
				ps, g := navigatePlot(start, []plant{}, garden)
				plots = append(plots, ps)
				garden = g
			}
		}
	}

	for _, plot := range plots {
		fmt.Println(plot)
		area := len(plot)
		perimiter := 0
		for _, plant := range plot {
			perimiter += 4 - plant.neighbors
		}
		price += area * perimiter
	}

	return price
}

func navigatePlot(start [2]int, plot []plant, garden [][]rune) ([]plant, [][]rune) {
	x, y := start[1], start[0]
	obj := plant{garden[y][x], x, y, 0}
	garden[y][x] = '.'
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, d := range directions {
		nextX, nextY := x+d[1], y+d[0]

		// This tile is out of bounds.
		if nextY < 0 || nextY >= len(garden) || nextX < 0 || nextX >= len(garden[0]) {
			continue
		}

		next := [2]int{nextY, nextX}
		if garden[nextY][nextX] == obj.kind {
			obj.neighbors += 1
			plot, garden = navigatePlot(next, plot, garden)
		}
	}

	plot = append(plot, obj)
	return plot, garden
}

// Parses the input data into a grid of runes.
func parseGarden(input string) [][]rune {
	rows := strings.Split(input, "\n")
	garden := make([][]rune, len(rows))
	for i, row := range rows {
		garden[i] = []rune(row)
	}
	return garden
}
