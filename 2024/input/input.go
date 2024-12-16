// Package input is for utilitize to handle input files.
package input

import (
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
