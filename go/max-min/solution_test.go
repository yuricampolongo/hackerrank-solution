package maxmin

import (
	"fmt"
	"testing"
)

func TestMaxMin(t *testing.T) {
	inputs := []struct {
		k        int32
		arr      []int32
		expected int32
	}{
		// {
		// 	k:        4,
		// 	arr:      []int32{1, 2, 3, 4, 10, 20, 30, 40, 100, 200},
		// 	expected: 3,
		// },
		// {
		// 	k:        2,
		// 	arr:      []int32{1, 2, 1, 2, 1},
		// 	expected: 0,
		// },
		// {
		// 	k:        3,
		// 	arr:      []int32{10, 100, 300, 200, 1000, 20, 30},
		// 	expected: 20,
		// },
		// {
		// 	k:        5,
		// 	arr:      []int32{4504, 1520, 5857, 4094, 4157, 3902, 822, 6643, 2422, 7288, 8245, 9948, 2822, 1784, 7802, 3142, 9739, 5629, 5413, 7232},
		// 	expected: 1335,
		// },
		{
			k:        3,
			arr:      []int32{100, 200, 300, 350, 400, 401, 402},
			expected: 2,
		},
	}

	for k, v := range inputs {
		t.Run(fmt.Sprintf("running test case #%d", k), func(t *testing.T) {
			result := maxMin(v.k, v.arr)

			if result != v.expected {
				t.Fatalf("expected %d got %d", v.expected, result)
			}
		})
	}
}
