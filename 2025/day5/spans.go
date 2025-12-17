package day5

// Span is a contiguous fresh sequence.
type Span struct {
	Start int
	End   int
}

func CombineRanges(spans []Span) []Span {
	// ... there's probably a better datastructure of this ...
	out := []Span{}
	for _, s := range spans {
		out = addSpan(out, s)
	}
	return out
}

func addSpan(spans []Span, toAdd Span) []Span {
	// none exist yet
	if len(spans) == 0 {
		return []Span{toAdd}
	}

	// before all ranges
	if toAdd.End < spans[0].Start {
		spans = append([]Span{toAdd}, spans...)
		return spans
	}

	// If it's in the range of the existing spans. Go span by span
	for idx, span := range spans {
		// if the toAdd span starts after the span ends, continue to the next, they're ordered
		if toAdd.Start > span.End {
			continue
		}

		if overlap(span, toAdd) {
			overlappingSpans := getAllOverlaps(toAdd, idx, spans)
			mergedSpan := mergeSpans(overlappingSpans)

			mergedList := append([]Span{}, spans[:idx]...)
			mergedList = append(mergedList, mergedSpan)
			mergedList = append(mergedList, spans[idx+len(overlappingSpans)-1:]...)
			return mergedList
		}

		// if the toAdd span ends before the start of the span, insert it before
		if toAdd.End < span.Start {
			result := append([]Span{}, spans[:idx]...)
			result = append(result, toAdd)
			result = append(result, spans[idx:]...)
			return result
		}
	}

	// remaining case is that it's  after all ranges
	return append(spans, toAdd)
}

// countFresh checks how many items fall in any of the spans, inclusively.
func countFresh(spans []Span, items []int) int {
	total := 0
	for _, item := range items {
		for _, span := range spans {
			if item >= span.Start && item <= span.End {
				total++
				break
			}
		}
	}
	return total
}

// overlap returns whether two spans touch.
func overlap(a, b Span) bool {
	// span a encompasses span b
	if a.Start <= b.Start && a.End >= b.End {
		return true
	}
	// span b encompasses span a
	if b.Start <= a.Start && b.End >= a.End {
		return true
	}
	// a and b overlap
	return a.Start <= b.Start && a.End >= b.Start || a.Start <= b.End && a.End >= b.End
}

// getAllOverlaps determines how many existing spans the toAdd span overlaps and returns their indicies.
func getAllOverlaps(test Span, idx int, existingSpans []Span) []Span {
	out := []Span{test}
	for idx <= len(existingSpans)-1 && overlap(test, existingSpans[idx]) {
		out = append(out, existingSpans[idx])
		idx++
	}
	return out
}

func mergeSpans(spans []Span) Span {
	minimumStart := spans[0].Start
	maximumEnd := spans[0].End
	for _, s := range spans {
		if s.Start < minimumStart {
			minimumStart = s.Start
		}
		if s.End > maximumEnd {
			maximumEnd = s.End
		}
	}
	return Span{
		Start: minimumStart,
		End:   maximumEnd,
	}
}

// countIDs takes ranges and counts how many total IDs are in all the ranges.
func countIDs(spans []Span) int {
	total := 0
	for _, span := range spans {
		total += (span.End - span.Start) + 1
	}
	return total
}
