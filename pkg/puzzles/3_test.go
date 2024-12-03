package puzzles

import "testing"

// Tests the full solution for day 3 based on the provided example.
func TestDayThree(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	eResult := 161

	if rResult, _ := dayThree(input); eResult != rResult {
		t.Errorf("DayThree() = %v, %v, want %v, %v", rResult, 0, eResult, 0)
	}
}
