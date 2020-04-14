package main

import(
	"io"
	"fmt"
)

func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}
