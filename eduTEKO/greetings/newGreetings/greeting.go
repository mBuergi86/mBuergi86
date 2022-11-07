package newgreetings

import (
	"fmt"
	"greeting-mod/hello"
)

func Greetings(name string) string {
	fmt.Scan(&name)

	message := fmt.Sprintln(hello.HelloPrint(name), "I'm glad you were here!")
	return message
}
