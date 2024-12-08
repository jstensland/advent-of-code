// Package runner has utilities for handling input files and day runs.
package runner

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Reader returns an io.Reader for the given file.
// If the file is not found, it fails here.
func Reader(inFile string) io.ReadCloser {
	in, err := os.Open(inFile) //nolint:gosec // Parser should protect against bad content
	if err != nil {
		log.Fatalf("failed to open file %s: %s", inFile, err)
	}
	return in
}

// TODO: add a function that Reads all for use with parsers that need
// the full file anyways

type Solver func(io.Reader) (int, error)

func RunIt(name string, fn Solver, inFile string) error {
	in := Reader(inFile)
	defer in.Close() //nolint:errcheck // no need to check for error

	answer, err := fn(in)
	if err != nil {
		return err
	}
	fmt.Printf("%s: %v\n", name, answer) //nolint:forbidigo // no IO CLI yet
	return nil
}
