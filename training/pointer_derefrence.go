package main

import "fmt"

func toChangeText(stringer *string) {
	*stringer = "Ich habe kein Geld für Wikipedia! 😢"
}

func main() {

	var stringPointer *string
	var text string

	text = "Willkommen bei Wikipeida!"
	stringPointer = &text

	fmt.Println(*stringPointer)

	*stringPointer = "Werde Mitglied bei Wikipedia!"

	fmt.Println(text)

	toChangeText(&text)

	fmt.Println(text)
}
