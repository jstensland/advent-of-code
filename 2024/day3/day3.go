package day3

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"text/scanner"
)

type Op struct {
	Left  int // would keep these private
	Right int // would keep these private
	// op string // always multiple for now
}

func RunPart1(in io.ReadCloser) (int, error) {
	// parse valid multiples
	ops := ParseOps(in)
	defer in.Close()

	return Compute(ops), nil
}

// ParseOps reads in the lines and parses each one, collecting operations
func ParseOps(in io.Reader) []Op {
	scanner := bufio.NewScanner(in)
	out := []Op{}
	for scanner.Scan() {
		ops, err := ParseLine(scanner.Text())
		if err != nil {
			log.Printf("error parsiing line: %v", err)
		}
		out = append(out, ops...)

	}
	if err := scanner.Err(); err != nil {
		log.Printf("error while scanning: %v", err)
		return nil
	}

	return out
}

func RunPart2(in io.ReadCloser) (int, error) {
	// parse valid multiples
	ops := ParseOps2(in)
	defer in.Close()

	return Compute(ops), nil
}

func Compute(ops []Op) int {
	total := 0
	for _, op := range ops {
		total += op.Left * op.Right
	}
	return total
}

// ParseOps reads in the lines and parses each one, collecting operations
func ParseOps2(in io.Reader) []Op {
	scanner := bufio.NewScanner(in)
	out := []Op{}
	// create parser out here so state can be maintened between lines
	parser := NewParser2()

	for scanner.Scan() {
		ops, err := parser.ParseLine(scanner.Text())
		if err != nil {
			log.Printf("error parsiing line: %v", err)
		}
		out = append(out, ops...)

	}
	if err := scanner.Err(); err != nil {
		log.Printf("error while scanning: %v", err)
		return nil
	}

	return out
}

func ParseOps3(in io.Reader) ([]Op, error) {
	// try using a scanner https://pkg.go.dev/text/scanner@go1.23.4
	// and checking tokens....
	//
	var s scanner.Scanner
	s.Init(in)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}
	return nil, nil
}

type opParser struct {
	src    string
	pos    int
	done   bool
	active bool
}

func (p opParser) Done() bool { return p.done }

func NewParser(in string) *opParser { return &opParser{src: in, active: true} }

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

// TODO: reduce need for those prints by breaking up seekOp
// into more testable parts

// seekOp finds the next op and returns it. If no op is found
// ok is false, and op value should be discarded.
func (p *opParser) seekOp() (Op, bool) {
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

type opParser2 struct {
	line   string
	pos    int
	done   bool
	active bool
}

func NewParser2() *opParser2 { return &opParser2{active: true} }

func (p opParser2) Done() bool { return p.done }

func (p *opParser2) ParseLine(in string) ([]Op, error) {
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
func (p *opParser2) seekOp2() (Op, bool) {
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
		p.pos = p.pos + toNextDo
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
