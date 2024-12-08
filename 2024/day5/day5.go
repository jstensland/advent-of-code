package day5

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/jstensland/advent-of-code/2024/runner"
)

func Run(inFile string) error {
	in := runner.Reader(inFile)
	defer in.Close()
	answer, err := SolvePart1(in) //nolint:forbidigo // no IO CLI yet
	if err != nil {
		return err
	}
	fmt.Println("Day 5 part 1:", answer)

	answer, err = SolvePart2(runner.Reader(inFile))
	if err != nil {
		return err
	}
	fmt.Println("Day 5 part 2:", answer)
	return nil
}

func SolvePart1(in io.Reader) (int, error) {
	rules, updates, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}

	total := 0
	for _, update := range updates {
		if rules.Check(update) {
			total += update.Middle()
		}
	}
	return total, nil
}

func SolvePart2(in io.Reader) (int, error) {
	rules, updates, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}

	total := 0
	for _, update := range updates {
		if rules.Check(update) {
			// skip if doesn't need a fix
			continue
		}
		slices.SortFunc(update, rules.SortFunc)
		total += update.Middle()
	}
	return total, nil
}

type Rules struct {
	// map of pages, to other pages that cannot come after
	exclusions map[int][]int
}

type Update []int

func (up Update) Middle() int {
	return up[len(up)/2]
}

func (r Rules) Check(up Update) bool {
	notLater := []int{}
	for _, page := range up {
		if slices.Contains(notLater, page) {
			return false
		}
		notLater = append(notLater, r.exclusions[page]...)
	}
	return true
}

// SortFunc compares the two pages and returns if a should come before b
// return a negative number when a < b
// positive number when a > b
// zero when a == b
//
// if there's a rule that a should come before b, return -1
func (r Rules) SortFunc(a, b int) int {
	// if 'a' can't come after 'b', b < a
	if slices.Contains(r.exclusions[a], b) {
		return 1
	}

	// if 'b' can't come after 'b', b < a
	if slices.Contains(r.exclusions[b], a) {
		return -1
	}

	return 0
}

// TODO: improve parsing...

func ParseInput(in io.Reader) (Rules, []Update, error) {
	rules := Rules{exclusions: map[int][]int{}}
	updates := []Update{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			// parse as a rule
			rawRule := strings.Split(line, "|")
			beforePage, err := strconv.Atoi(rawRule[0])
			if err != nil {
				return Rules{}, nil, fmt.Errorf("error parsing before page: %w", err)
			}
			afterPage, err := strconv.Atoi(rawRule[1])
			if err != nil {
				return Rules{}, nil, fmt.Errorf("error parsing after page: %w", err)
			}
			// check this number has been seen
			if _, ok := rules.exclusions[afterPage]; !ok {
				rules.exclusions[afterPage] = []int{}
			}
			rules.exclusions[afterPage] = append(rules.exclusions[afterPage], beforePage)

		} else if strings.Contains(line, ",") {
			update, err := parseUpdate(line)
			if err != nil {
				return Rules{}, nil, fmt.Errorf("error parsing update: %w", err)
			}
			updates = append(updates, update)
		}
	}

	return rules, updates, nil
}

func parseUpdate(in string) (Update, error) {
	rawUpdate := strings.Split(in, ",")
	out := Update{}
	for _, val := range rawUpdate {
		num, err := strconv.Atoi(val)
		if err != nil {
			return Update{}, fmt.Errorf("error parsing update number: %w", err)
		}
		out = append(out, num)
	}
	return out, nil
}
