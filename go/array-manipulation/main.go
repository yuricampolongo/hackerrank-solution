package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, Q int
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Scanf("%d %d", &N, &Q)
	sums := make([]int64, N+1)

	for i := 0; i < Q; i++ {
		var a, b, k int
		fmt.Scanf("%d %d %d", &a, &b, &k)
		sums[a] += int64(k)
		fmt.Printf("pos %d valor %d \n", a, sums[a])
		if b+1 <= N {
			sums[b+1] -= int64(k)
		}
	}

	fmt.Println(sums)

	/*matrix := [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}
	result := arrayManipulation(10, matrix)*/

}
