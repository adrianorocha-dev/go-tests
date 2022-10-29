package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Adriano")
	fmt.Println(message)
}
