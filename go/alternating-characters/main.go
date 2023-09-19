package main

import "fmt"

func main() {
	result := alternatingCharacters("AABAAB")
	fmt.Println(result)
}

func alternatingCharacters(s string) int32 {
	lastChar := ""
	deletions := 0

	for _, i := range s {
		letter := string(i)
		if lastChar == "" && letter != "" {
			lastChar = letter
		} else {
			if lastChar == letter {
				deletions++
			} else {
				lastChar = letter
			}
		}
	}

	return int32(deletions)
}
