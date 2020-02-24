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

	if language == spanish {
		return spanishhelloPrefix + name
	} else if language == french {
		return frenchhelloPrefix + name
	}

    return englishhelloPrefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}

