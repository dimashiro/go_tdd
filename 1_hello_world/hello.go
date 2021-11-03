package main

import "fmt"

const greetingMessage = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return greetingMessage + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
