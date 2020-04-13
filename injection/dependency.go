package main

import (
    "fmt"
    "os"
    "io"
)


// both bytes.Buffer and os.Stdout implement io.Writer
// So, io.Writer could be used both in test and application
func Greet(writer io.Writer, name string) {
	// fmt.Printf("Hello, %s", name) // send the string into the default stdout
	fmt.Fprintf(writer, "Hello, %s", name) // send the string into the buffer
}

func main() {
	Greet(os.Stdout, "Elodie")
}
