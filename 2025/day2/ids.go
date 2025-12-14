package day2

import (
	"errors"
	"fmt"
	"strconv"
)

type ID []int

func (id ID) AsInt() int {
	const base = 10
	result := 0
	for _, digit := range id {
		result = result*base + digit
	}
	return result
}

// TODO: I want this public but only able to be created via constructor..
// make a struct with private attribute id?
type InvalidID ID

var ErrIDNotInvalid = errors.New("the id was not invalid")

func NewInvalidID(id ID) (InvalidID, error) {
	if len(id)%2 != 0 {
		return nil, ErrIDNotInvalid
	}

	// If the ID isn't invalid, return an error
	for idx := range id[:len(id)/2] {
		if id[idx] != id[idx+len(id)/2] {
			return nil, ErrIDNotInvalid
		}
	}

	return InvalidID(id), nil
}

func (id InvalidID) NextInvalid() InvalidID {
	// we know we have an invalid ID, so it's even length

	// take the first half, turn it into an int
	firstHalf := ID(id[:len(id)/2]).AsInt()
	// increment it by one
	firstHalf++
	// turn it back into an invalid ID and return it
	whole := append(toID(firstHalf), toID(firstHalf)...)
	out, err := NewInvalidID(whole)
	if err != nil {
		panic(fmt.Errorf("finding next issue: %w", err)) // developer error
	}
	return out
}

func toID(in int) ID {
	str := strconv.Itoa(in)
	result := make(ID, 0, len(str))
	for _, digit := range str {
		result = append(result, int(digit-'0'))
	}
	return result
}

// func findFirst(id ID) InvalidID {
// 	return findFristRep(id, 2)
// }

func findFirst(id ID) InvalidID {
	var first ID
	startingLen := len(id)
	if startingLen%2 == 0 {
		firstHalf := id[:startingLen/2]
		// if the second half is bigger than the first, increment the first by 1 so the double concat is bigger
		// 1221 -> 1212 ❌
		// 1221 -> 1313 ✅
		// 183200 -> 184184✅
		// 565653 -> 566566 ✅
		secondHalf := id[startingLen/2:]
		if firstHalf.AsInt() < secondHalf.AsInt() {
			firstHalf = toID(firstHalf.AsInt() + 1)
		}
		first = append(firstHalf, firstHalf...)
	} else {
		// if start has an odd number of a digits, the lowest 1 followed by all zeros is the next candidate
		first = make(ID, startingLen+1)
		first[0] = 1
		first[(startingLen+1)/2] = 1
	}

	firstInvalid, err := NewInvalidID(first)
	if err != nil {
		panic(fmt.Errorf("finding first issue: %w", err))
	}

	if ID(firstInvalid).AsInt() < id.AsInt() {
		firstInvalid = firstInvalid.NextInvalid()
	}
	return firstInvalid
}

func findFirstV2(id ID) InvalidID {
	return findFristRep(id, 2)
}

func findFristRep(id ID, repSize int) InvalidID {
	// TODO: The search is now per n divisions, where n is the number of possible repeated digits
	// pick up here
	// - Make this algorithm a special case where n = 2, but also support n = 1, 3, 4, 5... etc.
	// - Only test for first number if evenly divisible length
	// - Consider how incrementing numbers push you to the else case
	// 899_999 -> 900_900
	// - consider cases for repeating needing to go up by 1. Longest repeat is best though if multiple.
	// e.g. 6 digit number can repeat 2 or 3
	// 122212 -> 123123 ✅
	// 122212 -> 131313 ❌
	//
	// algorithm
	// 1. determining the possible n
	// 1. try the longest 'n'. If it pushes you into the else, before increasing digits, try the next
	//    longest n
	// 1. if no n works, increase digits with largest n

	startingLen := len(id)
	toRepeat := []int{}
	if startingLen%repSize == 0 {
		firstPart := id[:startingLen/repSize]
		// if the second half is bigger than the first, increment the first by 1 so the double concat is bigger
		// 1221 -> 1212 ❌
		// 1221 -> 1313 ✅
		// 183200 -> 184184✅
		// 565653 -> 566566 ✅
		secondPart := id[startingLen/repSize : 2*startingLen/repSize]
		if firstPart.AsInt() < secondPart.AsInt() {
			toRepeat = toID(firstPart.AsInt() + 1)
		}
	} else {
		// if start has an odd number of a digits, the lowest 1 followed by all zeros is the next candidate
		toRepeat = make(ID, max(startingLen/repSize, 1))
		toRepeat[0] = 1
	}

	first := []int{}
	// repeat it
	for range repSize {
		first = append(first, toRepeat...)
	}

	firstInvalid, err := NewInvalidID(first)
	if err != nil {
		panic(fmt.Errorf("finding first issue: %w", err))
	}

	if ID(firstInvalid).AsInt() < id.AsInt() {
		firstInvalid = firstInvalid.NextInvalid()
	}
	fmt.Println("first invalid:", firstInvalid)
	return firstInvalid
}

func (id InvalidID) NextInvalidV2() InvalidID {
	// TODO: determine the algorithm to find the next invalid from a current one
	return InvalidID{}
}
