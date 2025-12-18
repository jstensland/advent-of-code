package day6

type Operator interface {
	Operate(operands []int) int
}

type Problem struct {
	operands []int
	operator Operator
}

// Operands returns a copy of the problem's operands.
func (p Problem) Operands() []int {
	out := make([]int, len(p.operands))
	copy(out, p.operands)
	return out
}

// Operator returns the problem's operator.
//
//nolint:ireturn // purposeful abstraction. We don't know the type of problem we have.
func (p Problem) Operator() Operator {
	return p.operator
}

//nolint:gochecknoglobals // sentinel value for addition
var Add add

type add struct{}

func (add) Operate(in []int) int {
	result := 0
	for _, num := range in {
		result += num
	}
	return result
}

//nolint:gochecknoglobals // sentinel value for multiplication
var Multiple multiple

type multiple struct{}

func (multiple) Operate(in []int) int {
	result := 1
	for _, num := range in {
		result *= num
	}
	return result
}

func (p Problem) Do() int {
	return p.operator.Operate(p.operands)
}

type Worksheet struct {
	problems []Problem
}

// Problems returns a copy of the worksheet problems.
func (w Worksheet) Problems() []Problem {
	out := make([]Problem, len(w.problems))
	copy(out, w.problems)
	return out
}

func (w Worksheet) Solve() []int {
	result := make([]int, len(w.problems))
	for idx, problem := range w.problems {
		result[idx] = problem.Do()
	}
	return result
}
