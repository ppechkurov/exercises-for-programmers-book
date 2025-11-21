package main

import (
	"exercices/madlib"
	"fmt"
	"os"
)

func main() {
	exitCode := 0
	if err := madlib.Main(os.Stdin, os.Stdout); err != nil {
		exitCode = 1
		fmt.Println(err)
	}
	os.Exit(exitCode)
}
