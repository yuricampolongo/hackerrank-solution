package triplesum

import (
	"sort"
)

func triplets(a []int32, b []int32, c []int32) int64 {
	a = removeDuplicates(a)
	b = removeDuplicates(b)
	c = removeDuplicates(c)

	sortInput(a)
	sortInput(b)
	sortInput(c)

	total := int64(0)

	for q := 0; q < len(b); q++ {
		qGreaterEqualA := findFirstIndexLessThanOrEqualTo(a, b[q])
		qGreaterEqualC := findFirstIndexLessThanOrEqualTo(c, b[q])

		if qGreaterEqualA == -1 || qGreaterEqualC == -1 {
			continue
		}

		total += int64((qGreaterEqualA + 1) * (qGreaterEqualC + 1))
	}

	return total
}

func findFirstIndexLessThanOrEqualTo(vals []int32, target int32) int {
	for i := len(vals) - 1; i >= 0; i-- {
		if vals[i] <= target {
			return i
		}
	}
	return -1
}

func sortInput(vals []int32) {
	sort.Slice(vals, func(i, j int) bool {
		return vals[i] < vals[j]
	})
}

func removeDuplicates(s []int32) []int32 {
	if len(s) < 1 {
		return s
	}

	// sort
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	prev := 1
	for curr := 1; curr < len(s); curr++ {
		if s[curr-1] != s[curr] {
			s[prev] = s[curr]
			prev++
		}
	}

	return s[:prev]
}
