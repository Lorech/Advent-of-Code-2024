package file

import (
	"fmt"
	"os"
)

// Reads the contents of a file, cleaning up the final newline character, and
// returns the contents of the file as a string.
//
// Helps prevent having to check for an empty line at the end of the file,
// which keeps causing issues for me.
func ReadFile(filename string) (string, error) {
	data, error := os.ReadFile(filename)

	if error != nil {
		return "", error
	}

	return string(data[:len(data)-1]), nil
}

// Reads the contents of a file for a specific day from the default input file
// directory.
func ReadInfile(day int) (string, error) {
	return ReadFile(fmt.Sprintf("infiles/%d.txt", day))
}

// Reads the contents of a file for a specific day's example from the default
// input file directory.
func ReadTestFile(day int) (string, error) {
	return ReadFile(fmt.Sprintf("infiles/%d_test.txt", day))
}
