// Package runner has generic runner logic to handle experimenting with each day
package runner

import "io"

// Solver is what each day part solver will implement. The reader is for the input.
type Solver func(in io.Reader) (int, error)
