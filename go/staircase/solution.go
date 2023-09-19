package main

import (
	"fmt"
	"strings"
)

func main() {
	staircase(4)
}

func staircase(n int32) {
	for step := int32(1); step <= n; step++ {
		fmt.Printf("%"+fmt.Sprint(n)+"s\n", strings.Repeat("#", int(step)))
	}
}
