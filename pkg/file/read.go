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
//
// Expects files to be stored in the "infiles" directory relative to the caller,
// with the filename format of "{day}.txt" or "{day}-{variation}.txt".
func ReadInfile(day int, variation ...int) (string, error) {
	if len(variation) > 0 {
		return ReadFile(fmt.Sprintf("infiles/%d-%d.txt", day, variation[0]))
	}

	return ReadFile(fmt.Sprintf("infiles/%d.txt", day))
}

// Reads the contents of a file for a specific day's example from the default
// input file directory.
//
// Expects files to be stored in the "infiles" directory relative to the caller,
// with the filename format of "{day}_test.txt" or "{day}_test-{variation}.txt".
func ReadTestFile(day int, variation ...int) (string, error) {
	if len(variation) > 0 {
		return ReadFile(fmt.Sprintf("infiles/%d_test-%d.txt", day, variation[0]))
	}

	return ReadFile(fmt.Sprintf("infiles/%d_test.txt", day))
}
