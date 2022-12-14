package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return englishHelloPrefix + name
}

func main() {
	test := "你好"
	fmt.Println("Hello, world" + test)
}
