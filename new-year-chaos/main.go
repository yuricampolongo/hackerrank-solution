package main

import (
	"fmt"
)

func main() {
	minimumBribes([]int32{1, 2, 5, 3, 7, 8, 6, 4})
}

func minimumBribes(q []int32) {
	n := len(q)
	chaos := false
	bribes := 0
	for i := n - 1; i >= 0; i-- {
		if q[i]-(int32(i)+1) > 2 {
			chaos = true
			break
		}
		index := int32(0)
		if q[i]-2 > 0 {
			index = q[i] - int32(2)
		}
		for j := index; j < int32(i); j++ {
			if q[j] > q[i] {
				bribes++
			}
		}
	}
	if chaos {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(bribes)
	}
}
