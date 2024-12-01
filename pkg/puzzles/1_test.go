package puzzles

import "testing"

// Tests the full solution for day 1 based on the provided example.
func TestDayOne(t *testing.T) {
	want := "Hello, World!"
	if got := DayOne(); got != want {
		t.Errorf("DayOne() = %v, want %v", got, want)
	}
}
