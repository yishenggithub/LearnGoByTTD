package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	test := "你好"
	fmt.Println("Hello, world" + test)
}
