package driver

import (
	"exercices/internal/input"
	"fmt"
	"os"
)

func Main() {
	isTerm := input.IsTTY(os.Stdin)
	if isTerm {
		fmt.Print("What is your age? ")
	}

	age := 0
	_, err := fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Printf("invalid age: %s", err)
		os.Exit(1)
	}

	msg := "You are "
	if age < 16 {
		msg += "not "
	}

	fmt.Println(msg + "old enough to legally drive")
}
