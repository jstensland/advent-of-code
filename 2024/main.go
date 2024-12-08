// Package main is a minimal entry point for running each day to help
// us get out of main and allow more testablility
package main

import (
	"log"

	"github.com/jstensland/advent-of-code/2024/runner"
)

// IMPROVEMENT: add IO/CLI handling
func main() {
	err := runner.Run()
	if err != nil {
		log.Fatal(err)
	}
}
