package main

import "fmt"

func main() {
	result := simpleArraySum([]int32{1, 2, 3, 4, 10, 11})
	fmt.Println(result)
}

func simpleArraySum(ar []int32) int32 {
	result := int32(0)
	for _, v := range ar {
		result += v
	}
	return int32(result)
}
