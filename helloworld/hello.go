package helloworld

const (
	englishHelloPrefix    = "Hello, "
	portugueseHelloPrefix = "Oieee, "
	frenchHelloPrefix     = "Bonjour, "
	french                = "FR"
	portuguese            = "PT-BR"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(lang)

	return prefix + name + "!"
}

func greetingPrefix(lang string) string {
	prefix := englishHelloPrefix
	switch lang {
	case "PT-BR":
		prefix = portugueseHelloPrefix
	case "FR":
		prefix = frenchHelloPrefix
	}

	return prefix
}

func main() {
}
