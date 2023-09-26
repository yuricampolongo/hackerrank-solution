package pairs

import "sort"

func pairs(k int32, arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	total := int32(0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]-arr[j] == k {
				total++
			} else if arr[i]-arr[j] > k {
				break
			}
		}
	}

	return total
}
