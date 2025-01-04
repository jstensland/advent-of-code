// Package day17 solves AoC 2024 day 17
package day17

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/jstensland/advent-of-code/2024/input"
)

func SolvePart1(in io.Reader) (string, error) {
	computer, err := ParseIn(in)
	if err != nil {
		return "", fmt.Errorf("error loading input: %w", err)
	}

	out := computer.RunProgram()
	return out, nil
}

type (
	Instruction func(*Computer, byte)
	OpCode      uint8 // this is really 0 to 7... use byte?
)

type Program struct {
	instructionIdx int // the instruction pointer to keep up to date as we process
	data           []uint8
}

// GetInstruction returns the instruction based on the opcode.
func (p *Program) GetInstruction(op OpCode) Instruction {
	ops := map[OpCode]Instruction{
		0: Adv,
		1: Bxl,
		2: Bst,
		3: Jnz,
		4: Bxc,
		5: Out,
		6: Bdv,
		7: Cdv,
	}
	fn, ok := ops[op]
	if !ok {
		panic(fmt.Sprintf("unhandled opcode %d", op))
	}
	return fn
}

// Next grabs the next value and increments the instruction pointer.
// Each op should call it twice to advance, except jump instructions.
func (p *Program) Next() uint8 {
	out := p.data[p.instructionIdx]
	p.instructionIdx++
	return out
}

type Computer struct {
	registerA int
	registerB int
	registerC int
	out       string
	Program   *Program
}

func (c *Computer) String() string {
	// IMPROVEMENT: output the "data" with * * around the current index spot, and drop index from output
	return fmt.Sprintf(`data:  %v
idx: %d
A: %d
B: %d
C: %d
out: %s`, c.Program.data, c.Program.instructionIdx, c.registerA, c.registerB, c.registerC, c.out)
}

func (c *Computer) RunProgram() string {
	for c.Program.instructionIdx < len(c.Program.data) {
		// fmt.Println("computer state looks like this:", c)
		code := OpCode(c.Program.Next())
		operand := c.Program.Next()
		c.Program.GetInstruction(code)(c, operand)
	}
	return c.Result()
}

// Result returns the accumulated 'out' but without the trailing comma
func (c *Computer) Result() string {
	return strings.Trim(c.out, ",")
}

//nolint:mnd // rules are magic
func (c *Computer) combo(operand byte) int {
	// Combo operands 0 through 3 represent literal values 0 through 3.
	switch int(operand) {
	case 0, 1, 2, 3:
		return int(operand)
	case 4:
		return c.registerA
	case 5:
		return c.registerB
	case 6:
		return c.registerC
	case 7:
		panic("reserved and should not appear")
	}
	panic("operand greater than 7 should be impossible")
}

// Adv instruction (opcode 0) performs division.
// The numerator is the value in the A register.
// The denominator is found by raising 2 to the power of the instruction's combo operand.
func Adv(c *Computer, operand byte) {
	c.registerA /= 1 << uint(c.combo(operand)) //nolint:gosec // guarding against values that are too large when parsing
}

// Bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's
// literal operand, then stores the result in register B.
func Bxl(c *Computer, operand byte) {
	c.registerB ^= int(operand)
}

// Bst instruction (opcode 2) calculates the value of its combo operand modulo 8
// (thereby keeping only its lowest 3 bits), then writes that value to the B register.
func Bst(c *Computer, operand byte) {
	// fmt.Println("combo operand", c.combo(operand))

	c.registerB = c.combo(operand) % 8 //nolint:mnd // magic computer in general!
}

// Jnz instruction (opcode 3) does nothing if the A register is 0. However, if the
// A register is not zero, it jumps by setting the instruction pointer to the value
// of its literal operand; if this instruction jumps, the instruction pointer is not
// increased by 2 after this instruction.
func Jnz(c *Computer, operand byte) {
	if c.registerA != 0 {
		c.Program.instructionIdx = int(operand)
	}
}

// Bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C,
// then stores the result in register B. (For legacy reasons, this instruction reads an
// operand but ignores it.)
func Bxc(c *Computer, _ byte) {
	c.registerB ^= c.registerC
}

// Out instruction (opcode 5) calculates the value of its combo operand modulo 8, then
// outputs that value. (If a program outputs multiple values, they are separated by commas.)
func Out(c *Computer, operand byte) {
	c.out += fmt.Sprintf("%d,", c.combo(operand)%8) //nolint:mnd // magic computer in general!
}

// Bdv instruction (opcode 6) works exactly like the adv instruction except that the result
// is stored in the B register. (The numerator is still read from the A register.)
func Bdv(c *Computer, operand byte) {
	c.registerB = c.registerA / (1 << uint(c.combo(operand))) //nolint:gosec // guarding against values that are too large
}

// Cdv instruction (opcode 7) works exactly like the adv instruction except that the result
// is stored in the C register. (The numerator is still read from the A register.)
func Cdv(c *Computer, operand byte) {
	c.registerC = c.registerA / (1 << uint(c.combo(operand))) //nolint:gosec // guarding against values that are too large
}

// ParseIn reads in all the games. For solving.
func ParseIn(in io.Reader) (*Computer, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(input.SplitOnDoubleCR)

	regA := regexp.MustCompile(`Register A: (\d+)`)
	regB := regexp.MustCompile(`Register B: (\d+)`)
	regC := regexp.MustCompile(`Register C: (\d+)`)
	program := regexp.MustCompile(`Program: (.+)`)

	var regAVal int
	var regBVal int
	var regCVal int
	var err error

	var data []uint8

	for scanner.Scan() {
		// read the first line as Button A
		line := scanner.Text()

		// If line starts with "Register" record the registers
		if strings.HasPrefix(line, "Register") {
			aM := regA.FindStringSubmatch(line)
			regAVal, err = strconv.Atoi(aM[1])
			if err != nil {
				return nil, fmt.Errorf("error parsing register A value: %w", err)
			}
			bM := regB.FindStringSubmatch(line)
			regBVal, err = strconv.Atoi(bM[1])
			if err != nil {
				return nil, fmt.Errorf("error parsing register B value: %w", err)
			}

			cM := regC.FindStringSubmatch(line)
			regCVal, err = strconv.Atoi(cM[1])
			if err != nil {
				return nil, fmt.Errorf("error parsing register C value: %w", err)
			}
			continue
		}
		// Otherwise, record the program
		pM := program.FindStringSubmatch(line)
		dataStr := strings.Split(pM[1], ",")
		for _, str := range dataStr {
			val, err := strconv.Atoi(str)
			if err != nil || val > 255 {
				return nil, fmt.Errorf("error parsing program data: %w", err)
			}
			data = append(data, uint8(val)) //nolint:gosec // guarding against values that are too large
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return NewComputer(regAVal, regBVal, regCVal, data), nil
}

func NewComputer(regA, regB, regC int, data []uint8) *Computer {
	return &Computer{
		registerA: regA,
		registerB: regB,
		registerC: regC,
		Program: &Program{
			instructionIdx: 0,
			data:           data,
		},
	}
}
