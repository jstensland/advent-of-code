package day2

import "slices"

// Levels returns a copy of the current report for inspection by tests
func (r Report) Levels() []int {
	return slices.Clone(r.levels)
}
