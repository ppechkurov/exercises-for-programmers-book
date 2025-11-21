package madlib

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

func prompt(what string, s *bufio.Scanner) (res string, err error) {
	s.Scan()
	res = strings.TrimSpace(s.Text())
	if res == "" {
		err = fmt.Errorf("%s is empty", what)
	}

	return res, err
}

func isTTY() bool {
	is := term.IsTerminal(int(os.Stdout.Fd()))
	fmt.Println("this is a tty")
	return is
}

func Main(in io.Reader, out io.Writer) error {
	s := bufio.NewScanner(in)

	inputs := []string{
		"a noun", "a verb", "an adjective", "an adverb",
	}

	results := []string{}

	for _, i := range inputs {
		fmt.Fprintf(out, "Enter %s: ", i)
		res, err := prompt(i, s)
		if err != nil {
			return err
		}

		if !isTTY() {
			fmt.Fprintln(out, res)
		}

		results = append(results, res)
		log.Printf("results: %v\n", results)
	}

	fmt.Fprintf(out,
		"Do you %s your %s %s %s? That's hilarious!",
		results[1], results[2], results[0], results[3])

	return nil
}
