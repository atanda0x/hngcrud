package main

import "fmt"

const englishPrefix = "Hello, "
const frenchPrefix = "Bonjour, "
const spanishPrefix = "Hola, "
const Spanish = "Spanish"
const French = "French"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == Spanish {
		return spanishPrefix + name
	}

	if language == French {
		return frenchPrefix + name
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("0x", "EnglIsh"))
}
