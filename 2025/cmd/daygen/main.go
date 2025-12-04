// Package main provides a generator for Advent of Code day boilerplate.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

const (
	defaultDirPerms  = 0o750
	defaultFilePerms = 0o640
)

type templateData struct {
	Day int
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var dayStr string
	flag.StringVar(&dayStr, "day", "", "Day number (1-25)")
	flag.Parse()

	if dayStr == "" {
		return errors.New("day number is required")
	}

	dayNum, err := strconv.Atoi(dayStr)
	if err != nil || dayNum < 1 || dayNum > 25 {
		return errors.New("invalid day number: must be between 1 and 25")
	}

	return generateDay(dayNum)
}

func generateDay(dayNum int) error {
	dayDir := fmt.Sprintf("day%d", dayNum)

	// Create directory
	if err := os.MkdirAll(dayDir, defaultDirPerms); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	data := templateData{Day: dayNum}

	// Generate day.go
	if err := generateFile(
		"cmd/daygen/templates/day.go.tmpl",
		filepath.Join(dayDir, fmt.Sprintf("day%d.go", dayNum)),
		data,
	); err != nil {
		return err
	}

	// Generate day_test.go
	if err := generateFile(
		"cmd/daygen/templates/day_test.go.tmpl",
		filepath.Join(dayDir, fmt.Sprintf("day%d_test.go", dayNum)),
		data,
	); err != nil {
		return err
	}

	//nolint:forbidigo // print is good enough here
	fmt.Printf("Generated boilerplate for day %d in %s/\n", dayNum, dayDir)
	return nil
}

func generateFile(templatePath, outputPath string, data templateData) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
	}

	//nolint:gosec // output paths are hard coded
	file, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, defaultFilePerms)
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("file %s already exists", outputPath)
		}
		return fmt.Errorf("failed to create file %s: %w", outputPath, err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
