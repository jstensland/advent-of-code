// Package main is a minimal entry point for running each day to help
// us get out of main and allow more testablility
package main

import (
	"advent2024/day3"
	"fmt"
	"log"
	"os"
)

// TODO: create a runner package that's invoked from here
// at most, this package should specify the days to run, similar
// to commandline arguments
func main() {
	// TODO: refactor day1 to allow running via tests
	// fmt.Println(day1.RunPart1("./day1/input.txt"))
	// fmt.Println(day1.RunPart2("./day1/input.txt"))

	// // day 2
	// answer, err := day2.RunPart1("./day2/input.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(answer)
	//
	// answer, err = day2.RunPart2("./day2/input.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(answer)

	// Day 3
	runDay3()
}

func runDay3() {
	inFile := "./day3/input.txt"
	in, err := os.Open(inFile)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to open file %s: %w", inFile, err))
		// return 0, fmt.Errorf("failed to open file %s: %w", inFile, err)
	}
	answer, err := day3.RunPart1(in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1", answer)

	inFile = "./day3/input.txt"
	in, err = os.Open(inFile)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to open file %s: %w", inFile, err))
		// return 0, fmt.Errorf("failed to open file %s: %w", inFile, err)
	}
	answer2, err := day3.RunPart2(in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 2", answer2)
}
