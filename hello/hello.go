package main

import "fmt"

const spanish = "Spanish"
const englishhelloPrefix = "Hello, "
const spanishhelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishhelloPrefix + name
	}

    return englishhelloPrefix + name
}

func main() {
    fmt.Println(Hello("world", ""))
}

