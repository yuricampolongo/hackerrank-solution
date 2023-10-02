package minimuncost

import (
	"fmt"
	"testing"
)

func TestMinimumCost(t *testing.T) {

	inputs := []struct {
		k        int32
		arr      []int32
		expected int32
	}{
		{
			k:        3,
			arr:      []int32{1, 2, 3, 4},
			expected: 11,
		},
		{
			k:        3,
			arr:      []int32{2, 5, 6},
			expected: 13,
		},
		{
			k:        2,
			arr:      []int32{2, 5, 6},
			expected: 15,
		},
		{
			k:        3,
			arr:      []int32{1, 3, 5, 7, 9},
			expected: 29,
		},
	}

	for k, v := range inputs {
		t.Run(fmt.Sprintf("running test case #%d", k), func(t *testing.T) {
			result := getMinimumCost(v.k, v.arr)

			if result != v.expected {
				t.Fatalf("expected %d got %d", v.expected, result)
			}
		})
	}
}
