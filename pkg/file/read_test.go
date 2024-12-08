package file

import "testing"

// Tests the reading and newline stripping of a file with a custom path.
func TestReadFile(t *testing.T) {
	r, err := ReadFile("infiles/1.txt")
	e := `1
2
3
4
5`

	if err != nil {
		t.Errorf("ReadFile() failed: %v", err)
		return
	}

	if r != e {
		t.Errorf("ReadFile() = %v, expected %v", r, e)
	}
}

// Tests the abstraction of reading an input file for a specific day.
func TestReadInfile(t *testing.T) {
	r, err := ReadInfile(1)
	e := `1
2
3
4
5`

	if err != nil {
		t.Errorf("ReadInfile() failed: %v", err)
		return
	}

	if r != e {
		t.Errorf("ReadInfile() = %v, expected %v", r, e)
	}
}

// Tests the abstraction of reading an input file with a variation for a specific day.
func TestReadInfileVariation(t *testing.T) {
	r, err := ReadInfile(1, 1)
	e := `6
7
8
9
10`

	if err != nil {
		t.Errorf("ReadInfile() failed: %v", err)
		return
	}

	if r != e {
		t.Errorf("ReadInfile() = %v, expected %v", r, e)
	}
}

// Tests the abstraction of reading a test input file for a specific day.
func TestReadTestFile(t *testing.T) {
	r, err := ReadTestFile(1)
	e := `1
2
3
4
5`

	if err != nil {
		t.Errorf("ReadTestFile() failed: %v", err)
		return
	}

	if r != e {
		t.Errorf("ReadTestFile() = %v, expected %v", r, e)
	}
}

// Tests the abstraction of reading a test input file for a specific day.
func TestReadTestFileWithVariation(t *testing.T) {
	r, err := ReadTestFile(1, 1)
	e := `6
7
8
9
10`

	if err != nil {
		t.Errorf("ReadTestFile() failed: %v", err)
		return
	}

	if r != e {
		t.Errorf("ReadTestFile() = %v, expected %v", r, e)
	}
}
