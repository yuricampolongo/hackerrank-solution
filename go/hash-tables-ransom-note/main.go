package main

import (
	"fmt"
	"sort"
)

func main() {
	/*reader := bufio.NewReader(os.Stdin)
	str, _, _ := reader.ReadLine()
	mn := strings.Split(string(str), " ")
	m, _ := strconv.Atoi(mn[0])
	n, _ := strconv.Atoi(mn[1])

	magazineText, _, _ := reader.ReadLine()
	magazineArray := make([]string, m)
	magazineInputs := strings.Split(string(magazineText), " ")
	magazineArray = append(magazineArray, magazineInputs...)

	noteText, _, _ := reader.ReadLine()
	noteArray := make([]string, n)
	noteInput := strings.Split(string(noteText), " ")
	noteArray = append(noteArray, noteInput...)

	checkMagazine(magazineArray, noteArray)*/

	//checkMagazine([]string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"}, []string{"ive", "got", "some", "coconuts"})
	//checkMagazine([]string{"two", "times", "three", "is", "not", "four"}, []string{"two", "times", "two", "is", "four"})
	//checkMagazine([]string{"all", "correct"}, []string{"all"})
	checkMagazine([]string{"all", "correct"}, []string{"All"})
}

func checkMagazine(magazine []string, note []string) {
	sort.Strings(magazine)
	for _, word := range note {
		posFound := sort.SearchStrings(magazine, word)
		if posFound >= len(magazine) {
			fmt.Println("No")
			return
		}
		if magazine[posFound] != word {
			fmt.Println("No")
			return
		}
		magazine = append(magazine[:posFound], magazine[posFound+1:]...)
	}
	fmt.Println("Yes")
}
