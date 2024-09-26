package main

import "fmt"

const (
	tagalogHelloPrefix = "Kamusta, "
	englishHelloPrefix = "Hello, "
	bisayaHelloPrefix  = "Maupay, "
	tagalog            = "Tagalog"
	bisaya             = "Bisaya"
	english            = "English"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	greetingMessage := greetingByLanguage(language) + name
	return greetingMessage
}

func greetingByLanguage(language string) (prefix string) {
	switch language {
	case tagalog:
		prefix = tagalogHelloPrefix
	case bisaya:
		prefix = bisayaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("PeterGun", "English"))
}
