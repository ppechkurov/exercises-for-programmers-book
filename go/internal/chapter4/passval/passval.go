package passval

import (
	"exercices/internal/input"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Main() {
	isTerm := input.IsTTY(os.Stdin)
	if isTerm {
		fmt.Print("What is your username? ")
	}

	username := ""
	_, err := fmt.Scanf("%s", &username)
	if err != nil {
		fmt.Println(fmt.Errorf("reading username: %w", err))
		os.Exit(2)
	}

	if isTerm {
		fmt.Print("What is the password? ")
	}
	password := ""
	_, err = fmt.Scanf("%s", &password)
	if err != nil {
		fmt.Println(fmt.Errorf("reading password: %w", err))
		os.Exit(2)
	}

	passwords := map[string][]byte{
		"leroy":  nil,
		"second": nil,
	}

	passwords["leroy"], _ = bcrypt.GenerateFromPassword([]byte("12345"), bcrypt.DefaultCost)
	passwords["second"], _ = bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	knownPassword, found := passwords[username]

	err = bcrypt.CompareHashAndPassword(knownPassword, []byte(password))
	if !found || err != nil {
		fmt.Println("I don't know you.")
		os.Exit(1)
	} else {
		fmt.Println("Welcome!")
	}
}
