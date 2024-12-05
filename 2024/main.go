// Package main is a minimal entry point for running each day to help
// us get out of main and allow more testablility
package main

import (
	"advent2024/day2"
	"fmt"
	"log"
)

func main() {
	// TODO: refactor day1 to allow running via tests
	// fmt.Println(day1.RunPart1("./day1/input.txt"))
	// fmt.Println(day1.RunPart2("./day1/input.txt"))

	// answer, err := day2.RunPart1("./day2/input.txt")
	answer, err := day2.RunPart2("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer)
}
