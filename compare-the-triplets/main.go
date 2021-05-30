package main

import "fmt"

func main() {
	result := compareTriplets([]int32{1, 2, 3}, []int32{3, 2, 1})
	fmt.Println(result)
}

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) []int32 {
	scores := []int32{0, 0}
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}

		if a[i] > b[i] {
			scores[0]++
		} else {
			scores[1]++
		}
	}
	return scores
}
