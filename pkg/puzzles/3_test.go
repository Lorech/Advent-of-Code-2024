package puzzles

import "testing"

var input string = `xmul(2,4)&mul[3,7]!^don't()do()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

// Tests the first part of the puzzle for day 3.
func TestDayThreePartOne(t *testing.T) {
	e := 161
	if r := d3p1(input); e != r {
		t.Errorf("d3p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 3.
func TestDayThreePartTwo(t *testing.T) {
	e := 48
	if r := d3p2(input); e != r {
		t.Errorf("d3p2() = %v, expected %v", r, e)
	}
}
