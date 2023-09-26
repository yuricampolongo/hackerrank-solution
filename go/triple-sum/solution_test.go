package triplesum

import (
	"fmt"
	"sort"
	"testing"
)

func Triplets(a []int32, b []int32, c []int32) int64 {
	// Sort the arrays to make it easier to find triplets
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })

	// Initialize counters and indices
	var count int64
	var i, j int

	// Loop through b array and find the number of elements in a and c that are less than or equal to each element in b
	for _, bb := range b {
		for i < len(a) && a[i] <= bb {
			i++
		}
		for j < len(c) && c[j] <= bb {
			j++
		}
		count += int64(i) * int64(j)
	}

	return count
}

func TestTriplets(t *testing.T) {
	testCases := []struct {
		a        []int32
		b        []int32
		c        []int32
		expected int64
	}{
		// {
		// 	[]int32{3, 5, 7},
		// 	[]int32{3, 6},
		// 	[]int32{4, 6, 9},
		// 	4,
		// },
		// {
		// 	[]int32{1, 4, 5},
		// 	[]int32{2, 3, 3},
		// 	[]int32{1, 2, 3},
		// 	5,
		// },
		{
			[]int32{1, 3, 5, 7},
			[]int32{5, 7, 9},
			[]int32{7, 9, 11, 13},
			12,
		},
	}

	for k, v := range testCases {
		t.Run(fmt.Sprintf("Running test case #%d", k), func(t *testing.T) {
			result := triplets(v.a, v.b, v.c)
			if result != v.expected {
				t.Fatalf("Expected %d got %d", v.expected, result)
			}
		})
	}
}
