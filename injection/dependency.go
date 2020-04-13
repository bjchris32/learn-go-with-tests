package main

import "bytes"
import "fmt"


func Greet(writer *bytes.Buffer, name string) {
	// fmt.Printf("Hello, %s", name) // send the string into the default stdout
	fmt.Fprintf(writer, "Hello, %s", name) // send the string into the buffer
}
