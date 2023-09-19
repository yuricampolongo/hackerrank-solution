package main

import "fmt"

func main() {
	result := aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005})
	fmt.Println(result)
}

/*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

func aVeryBigSum(ar []int64) int64 {
	result := int64(0)
	for _, v := range ar {
		result += v
	}
	return result
}
