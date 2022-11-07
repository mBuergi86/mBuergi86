package main

import "fmt"

func splitNums(word string, space []int) string {

	wordWithSpace := ""
	nums := 0

	for i := 0; i < len(space); i++ {
		for j := nums; j < space[i]; j++ {
			wordWithSpace += string([]rune(word)[j])
			nums += 1
		}
		wordWithSpace += " "
	}

	return wordWithSpace
}

func ibanEasySpace(word string) string {
	for i := 4; i < len(word); i += 5 {
		word = word[:i] + " " + word[i:]
	}
	return word
}

func main() {
	word := "ilovegeeksforgeeks"
	spaces := []int{1, 5, 10, 13, len(word)}

	fmt.Println(splitNums(word, spaces))

	iban := "CH3900700115651849173"

	fmt.Println(ibanEasySpace(iban))
}
