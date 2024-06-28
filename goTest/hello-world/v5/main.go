package main

const (
	French        = "French"
	spanish       = "Spanish"
	Yoruba        = "Yoruba"
	frenchPrefix  = "Bonjour, "
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	yorubaPrefix  = "Bawoni, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchPrefix
	case "Spanish":
		prefix = spanishPrefix
	case "Yoruba":
		prefix = yorubaPrefix
	default:
		prefix = englishPrefix
	}
	return

}
