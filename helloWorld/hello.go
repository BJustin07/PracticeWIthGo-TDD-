package helloWorld

const (
	tagalogHelloPrefix = "Kamusta, "
	englishHelloPrefix = "Hello, "
	bisayaHelloPrefix  = "Maupay, "
	tagalog            = "Tagalog"
	bisaya             = "Bisaya"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingByLanguage(language) + name
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
