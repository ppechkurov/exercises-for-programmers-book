package main

import (
	"exercices/sinterest"
	"fmt"
	"os"
)

func main() {
	err := sinterest.Main(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
