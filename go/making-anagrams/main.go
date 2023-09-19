package main

import "fmt"

func main() {
	result := makeAnagram("fcrxzwscanmligyxyvym", "jxwtrhvujlmrpdoqbisbwhmgpmeoke")
	fmt.Println(result)
}

func makeAnagram(a string, b string) int32 {
	original := make(map[string]int)
	deletions := 0

	for _, letter := range a {
		original[string(letter)]++
	}

	for _, l := range b {
		letter := string(l)
		if original[string(letter)] > 0 {
			original[string(letter)]--
		} else {
			deletions++
		}
	}

	for _, v := range original {
		if v > 0 {
			deletions += v
		}
	}

	return int32(deletions)
}
