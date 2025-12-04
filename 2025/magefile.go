//go:build mage

// Package main provides mage build targets.
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/magefile/mage/mg"
)

// Test runs all tests in the codebase.
func Test() error {
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, "go", "test", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("tests failed: %w", err)
	}
	return nil
}

// Lint runs golangci-lint on the codebase.
func Lint() error {
	mg.Deps(checkLinter)

	ctx := context.Background()
	cmd := exec.CommandContext(ctx, "golangci-lint", "run", "--build-tags", "mage", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("golangci-lint failed: %w", err)
	}
	return nil
}

var errLinterNotFound = errors.New(
	"golangci-lint not found. Install it with: " +
		"go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest",
)

// checkLinter verifies that golangci-lint is installed.
func checkLinter() error {
	if _, err := exec.LookPath("golangci-lint"); err != nil {
		return errLinterNotFound
	}
	return nil
}

const defaultDirPerms = 0o750

var (
	errNoSession  = errors.New("AOC_SESSION environment variable not set")
	errInvalidDay = errors.New("invalid day number: must be between 1 and 25")
	errBadStatus  = errors.New("failed to download input: bad status code")
)

// GetInput downloads the input for a specific day.
// Usage: mage getInput 1.
func GetInput(day string) error {
	dayNum, err := parseDay(day)
	if err != nil {
		return err
	}

	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return errNoSession
	}

	body, err := downloadInput(dayNum, session)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := body.Close(); closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()

	return saveInput(dayNum, body)
}

// NewDay generates boilerplate code for a new day.
// Usage: mage newDay 3.
func NewDay(day string) error {
	dayNum, err := parseDay(day)
	if err != nil {
		return err
	}

	ctx := context.Background()
	cmd := exec.CommandContext(ctx, "go", "run", "./cmd/daygen", "-day", strconv.Itoa(dayNum))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate day boilerplate: %w", err)
	}
	return nil
}

func parseDay(day string) (int, error) {
	dayNum, err := strconv.Atoi(day)
	if err != nil || dayNum < 1 || dayNum > 25 {
		return 0, errInvalidDay
	}
	return dayNum, nil
}

func downloadInput(dayNum int, session string) (io.ReadCloser, error) {
	url := fmt.Sprintf("https://adventofcode.com/2025/day/%d/input", dayNum)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to download input: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		if closeErr := resp.Body.Close(); closeErr != nil {
			return nil, errors.Join(fmt.Errorf("%w: %d", errBadStatus, resp.StatusCode), closeErr)
		}
		return nil, fmt.Errorf("%w: %d", errBadStatus, resp.StatusCode)
	}

	return resp.Body, nil
}

func saveInput(dayNum int, body io.Reader) (err error) {
	dayDir := fmt.Sprintf("day%d", dayNum)
	if err := os.MkdirAll(dayDir, defaultDirPerms); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	inputPath := filepath.Clean(filepath.Join(dayDir, "input.txt"))
	file, err := os.Create(inputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()

	if _, err := io.Copy(file, body); err != nil {
		return fmt.Errorf("failed to write input: %w", err)
	}

	return nil
}
