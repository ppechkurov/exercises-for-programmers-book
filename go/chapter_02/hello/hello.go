package hello

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

func Who(in io.Reader) (who string, err error) {
	s := bufio.NewScanner(in)
	s.Scan()
	who = s.Text()
	if who == "" {
		err = errors.New("input is empty")
	}
	return who, err
}

func Greet(who string) string {
	switch who {
	case "Peter":
		who = "Parker"
	case "Another":
		who = "Dawg"
	}

	return fmt.Sprintf("Hello, %s, nice to meet you!", who)
}

func Main(in io.Reader, out io.Writer) int {
	fmt.Fprintf(out, "What is your name? ")

	who, err := Who(in)
	if err != nil {
		fmt.Fprint(out, err)
		return 1
	}

	greet := Greet(who)
	fmt.Fprintln(out, greet)
	return 0
}
