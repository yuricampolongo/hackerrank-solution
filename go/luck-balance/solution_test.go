package luckbalance

import (
	"fmt"
	"testing"
)

func TestLuckBalance(t *testing.T) {
	inputs := []struct {
		k             int32
		contestsInput [][]int32
		expected      int32
	}{
		{
			k: 3,
			contestsInput: [][]int32{
				{5, 1}, {2, 1}, {1, 1}, {8, 1}, {10, 0}, {5, 0},
			},
			expected: 29,
		},
	}

	for k, v := range inputs {
		t.Run(fmt.Sprintf("Running test case #%d", k), func(t *testing.T) {
			result := luckBalance(v.k, v.contestsInput)

			if result != v.expected {
				t.Fatalf("expected %d got %d", v.expected, result)
			}
		})
	}
}
