package main

import (
	"fmt"
	"greeting-mod/newGreetings"
)

func main() {
	var name string
	fmt.Println("Please enter your name here: ")
	fmt.Println(newgreetings.Greetings(name))
}
