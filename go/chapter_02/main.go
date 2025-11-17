package main

import (
	"exercices/hello"
	"os"
)

func main() {
	os.Exit(hello.Main(os.Stdin, os.Stdout))
}
