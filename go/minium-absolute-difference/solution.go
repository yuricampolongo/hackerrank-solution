package miniumabsolutedifference

import "sort"

func minimumAbsoluteDifference(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	minimum := int32(-1)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			difference := arr[i] - arr[j]
			if difference < 0 {
				difference *= -1
			}

			if difference == 0 {
				return 0
			}

			if minimum == -1 || difference < minimum {
				minimum = difference
			} else {
				break
			}
		}
	}
	return minimum
}
