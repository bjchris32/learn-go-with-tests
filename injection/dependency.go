package main

import (
    "fmt"
    "os"
    "io"
    "net/http"
)


// both bytes.Buffer and os.Stdout implement io.Writer
// So, io.Writer could be used both in test and application
func Greet(writer io.Writer, name string) {
	// fmt.Printf("Hello, %s", name) // send the string into the default stdout
	fmt.Fprintf(writer, "Hello, %s", name) // send the string into the buffer
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// the First io.Writer in both os.Stdout and bytes.Buffer
	Greet(os.Stdout, "Elodie")
	// the second io.Writer in http.ResponseWriter
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
