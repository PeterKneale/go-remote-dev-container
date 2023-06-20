package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: mycli <input>\n")
		os.Exit(1)
	}

	input := os.Args[1]

	// your code here. This is a very simple example that just echoes the input.
	fmt.Printf("You input: %s\n", input)
}
