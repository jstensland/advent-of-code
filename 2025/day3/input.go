package day3

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ParseIn(r io.Reader) ([]Bank, error) {
	scanner := bufio.NewScanner(r)
	banks := []Bank{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		bank := make(Bank, 0, len(line))
		for _, char := range line {
			digit := int(char - '0')
			bank = append(bank, digit)
		}
		banks = append(banks, bank)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error on input: %w", err)
	}

	return banks, nil
}
