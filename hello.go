package main

import "fmt"

const greetingMessage = "Hello, "

func Hello(name string) string {
	return greetingMessage + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
