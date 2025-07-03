package main

import "fmt"

const (
	portugueseLang = "Portuguese"
	spanishLang    = "Spanish"
	frenchLang     = "French"

	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Ola, "
)

func main() {
	fmt.Println(Hello("Matheus", portugueseLang))
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

// lower case functions are private
func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanishLang:
		prefix = spanishHelloPrefix
	case frenchLang:
		prefix = frenchHelloPrefix
	case portugueseLang:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
