package hello

const (
	englishHelloPrefix = "Hello, "
	englishHelloSuffix = "World"
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(str string, language string) string {
	if len(str) == 0 {
		return englishHelloPrefix + englishHelloSuffix
	}

	if language == "Spanish" {
		return spanishHelloPrefix + str
	}

	if language == "French" {
		return frenchHelloPrefix + str
	}

	return englishHelloPrefix + str
}
