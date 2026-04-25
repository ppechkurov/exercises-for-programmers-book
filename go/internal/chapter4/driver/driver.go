package driver

import (
	"exercices/internal/input"
	"fmt"
	"os"
	"strings"
)

func Main() {
	isTerm := input.IsTTY(os.Stdin)
	if isTerm {
		fmt.Print("What is your age? ")
	}

	age := 0
	_, err := fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Printf("invalid age: %s\n", err)
		os.Exit(1)
	}
	if age < 0 {
		fmt.Println("invalid age: should be more than 0")
		os.Exit(1)
	}

	lookup := map[int][]string{
		16: {"US"},
		18: {"JP", "DE"},
	}

	msg := "You are "
	allowed, ok := lookup[age]
	if !ok {
		fmt.Println(msg + "not old enough to legally drive")
		os.Exit(2)
	}

	fmt.Printf(
		"%sold enough to legally drive in the following countries: %s\n",
		msg, strings.Join(allowed, ", "))
}
