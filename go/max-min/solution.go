package maxmin

import "sort"

func maxMin(k int32, arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	result := -1
	for i := 0; i <= len(arr)-int(k); i++ {
		subArr := arr[i : int(k)+i]
		diff := subArr[len(subArr)-1] - subArr[0]

		if diff == 0 {
			return 0
		} else if result == -1 || int(diff) < result {
			result = int(diff)
		}
	}

	return int32(result)
}
