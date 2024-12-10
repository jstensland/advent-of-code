// Package day9 solves AoC 2024 day 9.
package day9

import (
	"fmt"
	"io"
	"strconv"
)

func SolvePart1(in io.Reader) (int, error) {
	blocks, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}
	// fmt.Println(blocks)
	blocks.MoveFileSegments()
	// fmt.Println(blocks)
	return blocks.CheckSum(), nil
}

func SolvePart2(in io.Reader) (int, error) {
	blocks, err := ParseInput(in)
	if err != nil {
		return 0, fmt.Errorf("error parsing input file: %w", err)
	}
	// fmt.Println(blocks)
	blocks.MoveFiles()
	// fmt.Println(blocks)
	return blocks.CheckSum(), nil
}

type Blocks struct {
	raw        []int
	fileBlocks map[int]fileBlock
	largestID  int
}

func (b *Blocks) String() string {
	out := ""
	for _, v := range b.raw {
		if v == -1 {
			out += "."
		} else {
			out += strconv.Itoa(v)
		}
	}
	return out
}

// MoveFileSegments moves parts of files into empty spaces, and does not
// preserve file blocks
func (b *Blocks) MoveFileSegments() {
	// zero this out as this mapping is not maintained at all by this action
	b.fileBlocks = make(map[int]fileBlock)
	idxBackward := len(b.raw) - 1
	idxForward := 0
	var toMoveIdx int
	var emptyIdx int
	for {
		for {
			// find the last occupied space
			if b.raw[idxBackward] != -1 {
				toMoveIdx = idxBackward
				break
			}
			idxBackward-- // go backward through the blocks
		}

		for {
			// find the first unoccupied space
			if b.raw[idxForward] == -1 {
				emptyIdx = idxForward
				break
			}
			idxForward++ // go backward through the blocks
		}

		if idxForward >= idxBackward {
			break
		}
		b.Swap(toMoveIdx, emptyIdx)
	}
}

func (b *Blocks) MoveFiles() {
	// go through the files backward
	for i := b.largestID; i >= 0; i-- {
		// fmt.Println(b)
		fb := b.fileBlocks[i]

		idx := 0 // start at the beginning each time
		for {
			idx++
			if idx >= len(b.raw) {
				break // checked everywhere
			}
			// find unoccupied space
			if b.raw[idx] == -1 {
				// how much of it?
				startEmpty := idx
				for b.raw[idx] == -1 {
					idx++
					if idx >= len(b.raw) {
						break // hit the end
					}
				}
				endEmpty := idx
				// not moving right and big enough space?
				if endEmpty < fb.end && endEmpty-startEmpty >= fb.size() {
					b.SwapFile(fb, fileBlock{start: startEmpty, end: startEmpty + fb.size()})
					break
				}
			}
		}
	}
}

func (b *Blocks) Swap(i, j int) {
	b.raw[i], b.raw[j] = b.raw[j], b.raw[i]
}

func (b *Blocks) SwapFile(i, j fileBlock) {
	tmp := make([]int, i.size())
	copy(tmp, b.raw[i.start:i.end])
	copy(b.raw[i.start:i.end], b.raw[j.start:j.end])
	copy(b.raw[j.start:j.end], tmp)
}

func (b *Blocks) CheckSum() int {
	checkSum := 0
	for i, v := range b.raw {
		if v != -1 {
			checkSum += i * v
		}
	}
	return checkSum
}

type fileBlock struct {
	start int
	end   int
}

func (b fileBlock) size() int {
	return b.end - b.start
}

// ParseInput parses a long line into memory blocks, recording a mapping of files locations by ID
//
// ID each block and store info
// 12345 -> 0..111....22222
//
//	{0: {start: 0, end: 1}, 1: {start: 3, end: 6}, 2: {start: 10, end: 15}}
func ParseInput(in io.Reader) (*Blocks, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	raw := make([]int, 0)
	fm := make(map[int]fileBlock)
	id := 0
	for i, b := range data {
		if b == '\n' {
			continue
		}
		// since input is only 0..9, we can just convert the byte to an int
		val, err := strconv.Atoi(string(b))
		if err != nil {
			return nil, fmt.Errorf("failed to parse input byte: %w", err)
		}

		var store int // the value to store
		// every other number is space
		if i%2 == 0 {
			store = id
			id++
		} else {
			store = -1
		}

		start := len(raw)
		end := len(raw)
		for range val {
			raw = append(raw, store)
			end++
		}
		fm[store] = fileBlock{start: start, end: end}
	}

	return &Blocks{raw: raw, fileBlocks: fm, largestID: id}, nil
}
