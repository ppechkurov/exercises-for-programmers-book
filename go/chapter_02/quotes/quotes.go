package quotes

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/term"
)

func scan(s *bufio.Scanner) string {
	s.Scan()
	return strings.TrimSpace(s.Text())
}

func formatOut(author, quote string) string {
	return fmt.Sprintf("%s says, \"%s\"\n", author, quote)
}

func isTTY() bool {
	is := term.IsTerminal(int(os.Stdout.Fd()))
	fmt.Println("this is a tty")
	return is
}

func Main(in io.Reader, out io.Writer) error {
	fmt.Fprint(out, "What is the quote? ")
	s := bufio.NewScanner(in)
	quote := scan(s)
	if quote == "" {
		return errors.New("quote is empty")
	}

	if !isTTY() {
		fmt.Fprintln(out, quote)
	}

	fmt.Fprint(out, "Who said it? ")
	author := scan(s)
	if author == "" {
		return errors.New("author is empty")
	}

	if !isTTY() {
		fmt.Fprintln(out, author)
	}

	fmt.Fprint(out, formatOut(author, quote))
	return nil
}

type Quote struct {
	Quote  string
	Author string
}

func Challenge(qs []Quote, out io.Writer) {
	for _, q := range qs {
		fmt.Fprint(out, formatOut(q.Author, q.Quote))
	}
}
