package main

import (
	"exercices/count"
	"os"
)

func main() {
	os.Exit(count.Main(os.Stdin, os.Stdout))
}
