package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const arabicHelloPrefix = "أهلاً, "
const spanish = "Spanish"
const french = "French"
const arabic = "Arabic"

func main() {
	fmt.Println(hello("world", ""))
}

func hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case arabic:
		prefix = arabicHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
