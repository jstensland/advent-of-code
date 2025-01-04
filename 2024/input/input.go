// Package input is for utilitize to handle input files.
package input

import (
	"bytes"
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

// SplitOnDoubleCR implements Splitfunc for the scanner. https://pkg.go.dev/bufio#SplitFunc
// It will split the input on blank lines, that is, two carriage returns
func SplitOnDoubleCR(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil // should this return ErrFinalToken?
	}

	// Find the next double newline
	splitToken := []byte("\n\n")
	if i := bytes.Index(data, splitToken); i >= 0 {
		// We found a double carriage return.
		// Return the token up to that point.
		return i + len(splitToken), data[:i], nil
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}

	// Request more data.
	return 0, nil, nil
}
