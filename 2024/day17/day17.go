// Package day17 solves AoC 2024 day 17
package day17

import (
	"bufio"
	"fmt"
	"io"
	"iter"
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

// TODO: Start here... try producing fewer options. specifically, try powers of 2, or some close variation on that.
// need to determine which changes have any actual effect...
// Could try printing out the startin values before and after any actual change in the output to get a clearer idea

// Candidates
func Candidates() iter.Seq[int] {
	return func(yield func(int) bool) {
		// This produces numbers with trailing 0111110100
		i := 14836
		j := 15860
		for {
			if i < j {
				i += 16384
				// fmt.Println("yielding", i)
				if !yield(i) {
					return
				}
			} else {
				j += 16384
				// fmt.Println("yielding", j)
				if !yield(j) {
					return
				}
			}

			// but now another constraint for numbers like this...
			// 101001000100011010101011110111110100

			// 101001000000000000000000000000000000 = 44023414784

			// produce numbers with these bits set
			// 101001000000000000000000000000000000
			// AND last digits 0111110100
			// can safely do only numbers larger than this one.
			//
			//
			// e.g.

			// 		i := 0
			//
			// 		for {
			// 			// if i%1024 != 500 {
			// 			// looks like number has to end with 0111110100 == 500
			//
			// 			if i%16384 != 15860 && i%16384 != 14836 {
			// 				// 11100111110100 = 14836
			// 				// 11110111110100 = 15860
			// 				i++
			// 				continue
			// 			}
			//
			// 			fmt.Println("yielding", i)
			//
			// 			if !yield(i) {
			// 				return
			// 			}
			// 			i++
			//
			// 		}
			//     yielding 14836
			// yielding 15860
			// yielding 31220
			// yielding 32244
			// yielding 47604
			// yielding 48628
			// yielding 63988
			// yielding 65012
			// yielding 80372
			// yielding 81396
			// yielding 96756
			// yielding 97780
			// yielding 113140
			// yielding 114164
			// yielding 129524
			// yielding 130548
			// yielding 145908
			// yielding 146932
			// yielding 162292
			// yielding 163316
		}
	}
}

func SolvePart2BruteForce(in io.Reader) (int, error) {
	computer, err := ParseIn(in)
	if err != nil {
		return 0, fmt.Errorf("error loading input: %w", err)
	}

	// try printing starting values that successfully output the first value... look for patterns

	// i := 10000000
	// i = 151100000
	out := 0
	for i := range Candidates() {
		out = i

		// fmt.Println("trying:", i)

		// fmt.Println("data", computer.Program.DataString())

		//
		computer.Reset()
		computer.SetRegisterA(i)
		out := computer.RunProgram2(computer.Program.DataString())
		// out := computer.RunProgram()
		// fmt.Println("out", out)
		if out == computer.Program.DataString() {
			// fmt.Println("success!", i)
			// fmt.Println("out:", out)
			// fmt.Println("original:", computer.Program.DataString())
			break
		}

		if i%1_000_000_000 == 0 {
			fmt.Println("progress:", i)
		}

		// if i > 100_000_000_000 {
		// 	break // only do 10million for now
		// }
	}

	return out, nil
}

func SolvePart2Dynamic(in io.Reader) (int, error) {
	// TODO: idea is to try this with dynamic programming.
	// - recursive is one option
	// 	 - Check if you have the answer for the current state of the problem. Return if so
	// 	 - If not, compute the outcome for the current state, and save it before returning it
	// 	 - ... so what is likely being re-calculated? ... 🤔
	// 	 -
	// - ground up is the other
	//
	//
	// I just read through Reddit though, and it seems like it's more about noticing the pattern of the
	// program, and which values you can skip/should try...

	computer, err := ParseIn(in)
	if err != nil {
		return 0, fmt.Errorf("error loading input: %w", err)
	}

	out := 0
	for i := range Candidates() {
		out = i
		// fmt.Println("trying:", i)

		computer.Reset()
		computer.SetRegisterA(i)
		out := computer.RunProgram2(computer.Program.DataString())
		// out := computer.RunProgram()
		// fmt.Println("out", out)
		if out == computer.Program.DataString() {
			// fmt.Println("success!", i)
			// fmt.Println("out:", out)
			// fmt.Println("original:", computer.Program.DataString())
			break
		}

		if i%1_000_000_000 == 0 {
			fmt.Println("progress:", i)
		}

		// if i > 100_000_000_000 {
		// 	break // only do 10million for now
		// }
	}
	return out, nil
}

func SolvePart2LogicMyProgram(in io.Reader) (int, error) {
	_ = in
	// Program: 2,4,1,2,7,5,1,3,4,3,5,5,0,3,3,0

	// first action
	// - idx: 0
	// - code: 2 - Bst
	// - operand: Register A! - % 8 -> Register B: 4...     X % 8 -> Y
	//
	// - This implies differences for the last 8 bits of register A... only 8 though
	//
	// Second action
	// - idx: 2
	// - code: 1 -  Bxl
	// - operand: 2 -> Register B ^ 2 -> 4 XOR 2 = 6...   Y XOR 2
	//
	// so register B gets: X % 8 then XOR 2
	//
	// Third action
	// - idx: 4
	// - code: 7 - Cdv -> READ REGISTER A! / 2^0 = Register A -> Register C
	// - oprand 5 - Reg B
	//
	// register C gets: X / 2^(Reg B) here reg B is 0 to 7, so 1, 2, 4, 8, 16, 32, 64, 128
	// - and whole number result gets stored... so meaningful differences in register A
	//   would require a low % 8, or big jumps. if % 8 = 7, Reg C only changes every 128 X values

	// Fourth action
	// - idx: 6
	// - code: 1 - Bxl
	// - operand: 3 -> Register B ^ 3 -> Reg B
	//
	// so register B gets: X % 8 then XOR 2 then XOR 3 or X % 8 XOR 1

	// Fifth action
	// - idx: 8
	// - code: 4 - Bxc - Reg B XOR Reg C -> Reg B
	// - operand: 3 ignored.   (X % 8 XOR 1) XOR ( X / 2^Reg B), where Reg B is 0 to 7

	// Sixth action
	// - idx: 8
	// - code: 5 - Out
	// - operand: 5 -> Reg B % 8 gets output. MUST BE 2! as it's the first value

	// 2 = ((X % 8 XOR 1) XOR ( X / 2^Reg B), where Reg B is 0 to 7) % 8
	//
	// Reg B 0 to 7 means 1, 2, 4, 8, 16, 32, 64, 128
	// Reg B means

	return 0, nil
}

type (
	Instruction func(*Computer, byte)
	OpCode      uint8 // this is really 0 to 7... use byte?
)

type Program struct {
	instructionIdx int // the instruction pointer to keep up to date as we process
	data           []uint8
}

func (p *Program) DataString() string {
	output := ""
	for _, val := range p.data {
		output += strconv.Itoa(int(val)) + ","
	}
	return strings.Trim(output, ",")
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
	startingA int
	newOut    bool
	reset     func(*Computer)
	registerA int
	registerB int
	registerC int
	out       string
	Program   *Program
}

func NewComputer(regA, regB, regC int, data []uint8) *Computer {
	return &Computer{
		reset: func(c *Computer) {
			c.registerA = regA
			c.registerB = regB
			c.registerC = regC
			c.out = ""
			c.Program.instructionIdx = 0
		},
		registerA: regA,
		registerB: regB,
		registerC: regC,
		Program: &Program{
			instructionIdx: 0,
			data:           data,
		},
	}
}

func (c *Computer) Reset() {
	c.reset(c)
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

func (c *Computer) SetRegisterA(in int) {
	c.startingA = in
	c.registerA = in
}

func (c *Computer) RunProgram() string {
	for c.Program.instructionIdx < len(c.Program.data) {
		fmt.Println("computer state looks like this:", c)
		code := OpCode(c.Program.Next())
		operand := c.Program.Next()
		c.Program.GetInstruction(code)(c, operand)
	}
	return c.Result()
}

func (c *Computer) RunProgram3(_ string) string {
	// check if we're done
	for c.Program.instructionIdx < len(c.Program.data) {
		// create a map key that includes all the information about the current
		// starting state
		//
		// compute the result of that starting state, and store it
		// - result is anythig output and where to jump to
		//
		//
		fmt.Println("computer state looks like this:", c)
		code := OpCode(c.Program.Next())
		operand := c.Program.Next()
		c.Program.GetInstruction(code)(c, operand)
	}
	return c.Result()
}

func (c *Computer) RunProgram2(answer string) string {
	for c.Program.instructionIdx < len(c.Program.data) {
		// fmt.Println("computer state looks like this:", c)

		code := OpCode(c.Program.Next())
		operand := c.Program.Next()
		c.Program.GetInstruction(code)(c, operand)
		if !strings.HasPrefix(answer, c.Result()) {
			return "fail"
			// } else if c.newOut && c.Result() != "" && len(c.Result()) > 15 {
		} else if c.newOut && c.Result() != "" && len(c.Result()) > 19 {
			// 			fmt.Printf(`startingA: %b
			// startingA mod 8: %v
			// answer so far: %v
			// `, c.startingA, c.startingA%8, c.Result())

			fmt.Printf(`startingA: %64b - %16d - %v - answer so far: %v
`, c.startingA, c.startingA, c.startingA%16384, c.Result())

			// experimenting with 3rd value...
			//
			// Try subtracting powers of 2 from the decimal value to see if it's consistent offset/jumps
			//
			// fmt.Println("computer state looks like this:", c)
			c.newOut = false
		}
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
	c.newOut = true
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
