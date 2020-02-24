package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishhelloPrefix = "Hello, "
const spanishhelloPrefix = "Hola, "
const frenchhelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishhelloPrefix
	switch language {
	case french:
		prefix = frenchhelloPrefix
	case spanish:
		prefix = spanishhelloPrefix
	}

    return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchhelloPrefix
	case spanish:
		prefix = spanishhelloPrefix
	default:
		prefix = englishhelloPrefix
	}
	return
}

func main() {
    fmt.Println(Hello("world", ""))
}

