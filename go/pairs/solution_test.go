package pairs

import (
	"fmt"
	"testing"
)

func TestPairs(t *testing.T) {
	testCases := []struct {
		arr      []int32
		k        int32
		expected int32
	}{
		{[]int32{1, 5, 3, 4, 2}, 2, 3},
		{[]int32{1, 2, 3, 4}, 1, 3},
	}

	for k, v := range testCases {
		t.Run(fmt.Sprintf("Testing input %d", k), func(t *testing.T) {
			result := pairs(v.k, v.arr)
			if result != v.expected {
				t.Errorf("Expected result %d but got %d", v.expected, result)
			}
		})
	}

}
