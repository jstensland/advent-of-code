// Package day3 contains the solution to the Advent of Code 2024 Day 3 puzzle.
package day3

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Op struct {
	Left  int // would keep these private
	Right int // would keep these private
	// op string // always multiple for now
}

func SolvePart1(in io.Reader) (int, error) {
	return Compute(ParseOps(in)), nil
}

// ParseOps reads in the lines and parses each one, collecting operations
func ParseOps(in io.Reader) []Op {
	inScanner := bufio.NewScanner(in)
	out := []Op{}
	for inScanner.Scan() {
		ops, err := ParseLine(inScanner.Text())
		if err != nil {
			log.Printf("error parsiing line: %v", err)
		}
		out = append(out, ops...)
	}
	if err := inScanner.Err(); err != nil {
		log.Printf("error while scanning: %v", err)
		return nil
	}

	return out
}

func SolvePart2(in io.Reader) (int, error) {
	return Compute(ParseOps2(in)), nil
}

func Compute(ops []Op) int {
	total := 0
	for _, op := range ops {
		total += op.Left * op.Right
	}
	return total
}

// ParseOps2 reads in the lines and parses each one, collecting operations
func ParseOps2(in io.Reader) []Op {
	inScanner := bufio.NewScanner(in)
	out := []Op{}
	// create parser out here so state can be maintened between lines
	parser := NewParser2()

	for inScanner.Scan() {
		ops, err := parser.ParseLine(inScanner.Text())
		if err != nil {
			log.Printf("error parsiing line: %v", err)
		}
		out = append(out, ops...)
	}
	if err := inScanner.Err(); err != nil {
		log.Printf("error while scanning: %v", err)
		return nil
	}

	return out
}

type OpParser struct {
	src    string
	pos    int
	done   bool
	active bool
}

func (p *OpParser) Done() bool { return p.done }

func NewParser(in string) *OpParser { return &OpParser{src: in, active: true} }

func ParseLine(in string) ([]Op, error) {
	parser := NewParser(in)
	out := []Op{}
	for !parser.Done() {
		op, ok := parser.seekOp()
		if ok {
			out = append(out, op)
		}
	}
	return out, nil
}

// IMPROVEMENT: reduce need for those prints by breaking up seekOp
// into more testable parts

// seekOp finds the next op and returns it. If no op is found
// ok is false, and op value should be discarded.
func (p *OpParser) seekOp() (Op, bool) {
	// read down the line from the current position until you find 'mul('
	// fmt.Println("searching for mul in", p.src[p.pos:])
	toNextMul := strings.Index(p.src[p.pos:], "mul(")
	if toNextMul < 0 {
		// no ops left
		p.done = true
		return Op{}, false
	}

	mulIdx := p.pos + toNextMul
	// fmt.Println("mul( index", mulIdx)

	// increment the position until after this `mul(`
	p.pos = mulIdx + len("mul(")
	// fmt.Println("new position", p.pos)

	// search for the next comma
	commaIdx := p.pos + strings.Index(p.src[p.pos:], ",")
	// fmt.Println("comma index", commaIdx)
	leftVal, err := strconv.Atoi(p.src[p.pos:commaIdx])
	if err != nil {
		return p.seekOp() // try again
	}
	// fmt.Println("left value", leftVal)

	// search for the next `)`
	endParenIdx := p.pos + // current seeking position after mul(
		(commaIdx - p.pos) + // length of first and a comma
		strings.Index(p.src[commaIdx:], ")") // distance to the end paren
	// fmt.Println("end paren index", endParenIdx)

	// try to convert right value
	rightVal, err := strconv.Atoi(p.src[commaIdx+1 : endParenIdx])
	if err != nil {
		return p.seekOp() // try again
	}

	return Op{Left: leftVal, Right: rightVal}, true
}

// IMPROVEMENT: consolidate with OpParser

type OpParser2 struct {
	line   string
	pos    int
	done   bool
	active bool
}

func NewParser2() *OpParser2 { return &OpParser2{active: true} }

func (p *OpParser2) Done() bool { return p.done }

func (p *OpParser2) ParseLine(in string) ([]Op, error) {
	p.line = in    // new line
	p.pos = 0      // start at the beginning
	p.done = false // not done before we start!
	out := []Op{}
	for !p.Done() {
		op, ok := p.seekOp2()
		if ok {
			out = append(out, op)
		}
	}
	// fmt.Println("line:", in, "made ops:", out)
	return out, nil
}

// seekOp finds the next op and returns it. If no op is found
// ok is false, and op value should be discarded.
func (p *OpParser2) seekOp2() (Op, bool) {
	// fmt.Println("start seekOp2")

	if !p.active {
		// progress to the next do()
		// fmt.Println("searching for do() in", p.line[p.pos:])
		toNextDo := strings.Index(p.line[p.pos:], "do()") // find the next do
		// fmt.Println("toNextDo is:", toNextDo)
		if toNextDo < 0 {
			// no do() left on the line. keep inactive state and go to the next one
			p.done = true
			return Op{}, false
		}
		p.pos += toNextDo
		p.active = true
	}
	// always active below here

	// search for the next `don't()`
	// fmt.Println("searching for don't() in", p.line[p.pos:])
	toNextDont := strings.Index(p.line[p.pos:], "don't()")
	// fmt.Println("toNextDont is:", toNextDont)

	// read down the line from the current position until you find 'mul('
	// fmt.Println("searching for mul in", p.line[p.pos:])
	toNextMul := strings.Index(p.line[p.pos:], "mul(")
	// fmt.Println("toNextMul is:", toNextMul)
	if toNextMul < 0 && toNextDont < 0 {
		// no more mul or don't, and we're already active. go to the next line
		p.done = true
		return Op{}, false
	}

	// if `don't()` comes next
	if toNextDont > 0 && toNextDont < toNextMul {
		// set the parser to inactive
		p.active = false
		// progress to after this don't() and try again
		p.pos = p.pos + toNextDont + len("don't()")
		return p.seekOp2()
	}

	// if no more `mul(` go to next line
	if toNextMul < 0 {
		return Op{}, false
	}

	mulIdx := p.pos + toNextMul
	// fmt.Println("mul( index", mulIdx)

	// increment the position until after this `mul(`
	p.pos = mulIdx + len("mul(")
	// fmt.Println("new position", p.pos)

	// search for the next comma
	commaIdx := p.pos + strings.Index(p.line[p.pos:], ",")
	// fmt.Println("comma index", commaIdx)
	leftVal, err := strconv.Atoi(p.line[p.pos:commaIdx])
	if err != nil {
		return p.seekOp2() // try again
	}
	// fmt.Println("left value", leftVal)

	// search for the next `)`
	endParenIdx := p.pos + // current seeking position after mul(
		(commaIdx - p.pos) + // length of first and a comma
		strings.Index(p.line[commaIdx:], ")") // distance to the end paren
	// fmt.Println("end paren index", endParenIdx)

	// try to convert right value
	rightVal, err := strconv.Atoi(p.line[commaIdx+1 : endParenIdx])
	if err != nil {
		return p.seekOp2() // try again
	}

	return Op{Left: leftVal, Right: rightVal}, true
}
