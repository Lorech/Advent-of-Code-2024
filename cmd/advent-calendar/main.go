package main

import (
	"fmt"
	"lorech/advent-of-code-2024/pkg/puzzles"
	"os"
)

func main() {
	solved := 7
	for day := 1; day <= solved; day++ {
		filename := fmt.Sprintf("infiles/%d.txt", day)
		data, error := os.ReadFile(filename)
		if error != nil {
			panic(error)
		}

		one, two, err := puzzles.Solve(day, string(data))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Day %v: %d, %d\n", day, one, two)
		}
	}
}
