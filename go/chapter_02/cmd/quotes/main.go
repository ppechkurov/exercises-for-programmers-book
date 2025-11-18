package main

import (
	"exercices/quotes"
	"fmt"
	"os"
)

func main() {
	exitCode := 0
	if err := quotes.Main(os.Stdin, os.Stdout); err != nil {
		exitCode = 1
		fmt.Println(err)
	}
	os.Exit(exitCode)
}
