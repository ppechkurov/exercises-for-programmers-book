package main

import (
	"exercices/internal/chapter4/taxcalc"
	"fmt"
	"os"
)

func main() {
	err := taxcalc.Main()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
