package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 2.
func TestDayTwoPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 2
	if r := d2p1(input); e != r {
		t.Errorf("d2p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 2.
func TestDayTwoPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(2)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 4
	if r := d2p2(input); e != r {
		t.Errorf("d2p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 2.
func BenchmarkDayTwoPartOne(b *testing.B) {
	input, err := file.ReadInfile(2)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d2p1(input)
	}
}

// Benchmarks the second part of the puzzle for day 2.
func BenchmarkDayTwoPartTwo(b *testing.B) {
	input, err := file.ReadInfile(2)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d2p2(input)
	}
}
