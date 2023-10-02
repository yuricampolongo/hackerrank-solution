package miniumabsolutedifference

import (
	"fmt"
	"testing"
)

func TestMinimiumDifference(t *testing.T) {
	inputs := []struct {
		arr      []int32
		expected int32
	}{
		{
			arr:      []int32{3, -7, 0},
			expected: 3,
		},
		{
			arr:      []int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75},
			expected: 1,
		},
	}

	for k, v := range inputs {
		t.Run(fmt.Sprintf("Running test case #%d", k), func(t *testing.T) {
			result := minimumAbsoluteDifference(v.arr)

			if result != v.expected {
				t.Fatalf("Expected %d got %d", v.expected, result)
			}
		})
	}

}
