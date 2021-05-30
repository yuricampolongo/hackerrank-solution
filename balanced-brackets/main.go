package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isBalanced(s string) string {
	stack := make([]string, len(s))
	opens, _ := regexp.Compile(`\(|\[|\{`)
	totalOpens := 0

	for _, i := range s {
		char := string(i)
		if opens.Match([]byte(char)) {
			stack = append(stack, char)
			totalOpens++
		} else {
			if totalOpens == 0 {
				return "NO"
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			totalOpens--
			switch pop {
			case "{":
				if char != "}" {
					return "NO"
				}
			case "(":
				if char != ")" {
					return "NO"
				}
			case "[":
				if char != "]" {
					return "NO"
				}
			}
		}
	}
	if totalOpens > 0 {
		return "NO"
	}

	return "YES"

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		result := isBalanced(s)

		fmt.Printf("\n%s\n", result)
	}

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
