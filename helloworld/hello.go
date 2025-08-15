package main

const (
	ENGLISHGREETINGPREFIX = "Hello, "
	SPANISHGREETINGPREFIX = "Hola, "
	FRENCHGREETINGPREFIX  = "Bonjour, "
	spanish               = "Spanish"
	french                = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World!"
	}

	return greeting(language) + name
}

func greeting(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = SPANISHGREETINGPREFIX
	case french:
		prefix = FRENCHGREETINGPREFIX
	default:
		prefix = ENGLISHGREETINGPREFIX
	}
	return
}
