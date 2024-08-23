package main

const (
	helloPrefix        = "hello "
	spanishHelloPrefix = "Hola "
	frenchHelloPrefix  = "Bonjour, "
	spanish            = "Spanish"
	french             = "French"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}
