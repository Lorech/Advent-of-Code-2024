package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first example of the puzzle for day 10 with one trailhead.
func TestDayTenPartOneFirstSample(t *testing.T) {
	input, err := file.ReadTestFile(10, "1")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 1
	if r := d10p1(input); e != r {
		t.Errorf("d10p1() = %v, expected %v", r, e)
	}
}

// Tests the second example of the puzzle for day 10 with the trail in an
// upside-down Y shape.
func TestDayTenPartOneSecondSample(t *testing.T) {
	input, err := file.ReadTestFile(10, "2")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 2
	if r := d10p1(input); e != r {
		t.Errorf("d10p1() = %v, expected %v", r, e)
	}
}

// Tests the third example of the puzzle for day 10 with the wider trail moving
// in many directions relative to the trailhead.
func TestDayTenPartOneThirdSample(t *testing.T) {
	input, err := file.ReadTestFile(10, "3")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 4
	if r := d10p1(input); e != r {
		t.Errorf("d10p1() = %v, expected %v", r, e)
	}
}

// Tests the second example of the puzzle for day 10 with two trails
// intersecting while moving diagonally.
func TestDayTenPartOneFourthSample(t *testing.T) {
	input, err := file.ReadTestFile(10, "4")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 3
	if r := d10p1(input); e != r {
		t.Errorf("d10p1() = %v, expected %v", r, e)
	}
}

// Tests the first part of the puzzle for day 10.
func TestDayTenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(10)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 36
	if r := d10p1(input); e != r {
		t.Errorf("d10p1() = %v, expected %v", r, e)
	}
}
