package main

import "fmt"

func main() {
	result := rotLeft([]int32{1, 2, 3, 4, 5}, 2)
	fmt.Println(result)
}

/*
 * Complete the 'rotLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER d
 */

func rotLeft(a []int32, d int32) []int32 {
	if d == 0 {
		return a
	}

	pos := a[0]

	for i := 0; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
	a[len(a)-1] = pos
	return rotLeft(a, d-1)
}
