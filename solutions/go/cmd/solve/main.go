package main

import (
	"flag"
	"fmt"
	"lorech/go-advent-of-code/pkg/file"
	"lorech/go-advent-of-code/pkg/puzzles"
)

func main() {
	start, end := 1, 25

	pDay := flag.Int("day", -1, "Solve a specific day; solves all days by default")
	flag.Parse()

	if *pDay != -1 {
		start, end = *pDay, *pDay
	}

	for day := start; day <= end; day++ {
		data, error := file.ReadInfile(2024, day)
		if error != nil {
			panic(error)
		}

		one, two, err := puzzles.Solve(day, data)
		if err != nil {
			fmt.Printf("Day %d: %v\n", day, err)
		} else {
			fmt.Printf("Day %d: %v, %v\n", day, one, two)
		}
	}
}
