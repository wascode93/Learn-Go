package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(hello("world"))
}

func hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name
}
