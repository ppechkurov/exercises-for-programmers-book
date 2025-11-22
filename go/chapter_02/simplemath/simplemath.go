package simplemath

import (
	"bufio"
	"errors"
	"exercices/tty"
	"fmt"
	"io"
	"strconv"
)

var (
	ErrEmpty    = errors.New("input is empty")
	ErrInvalid  = errors.New("input invalid")
	ErrNegative = errors.New("negative input")
)

func getInput(s *bufio.Scanner) (res string, err error) {
	s.Scan()
	res = s.Text()
	if res == "" {
		err = ErrEmpty
	}
	return res, err
}

func Main(in io.Reader, out io.Writer) error {
	s := bufio.NewScanner(in)

	inputs := make([]int, 2)

	for i, v := range []string{"first", "second"} {
		fmt.Fprintf(out, "What is the %s number? ", v)
		input, err := getInput(s)
		if err != nil {
			return err
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			return fmt.Errorf("parsing input %q: %w", input, ErrInvalid)
		}
		if num < 0 {
			return fmt.Errorf("only positive numbers are allowed: %w", ErrNegative)
		}

		if !tty.IsTTY() {
			fmt.Fprintln(out, input)
		}
		inputs[i] = num
	}

	first, second := inputs[0], inputs[1]
	for _, op := range []string{"+", "-", "*", "/"} {
		fmt.Fprintf(out,
			"%d %s %d = %d\n",
			first, op, second, calc(op, first, second))
	}

	return nil
}

func calc(op string, first, second int) int {
	switch op {
	case "+":
		return first + second
	case "-":
		return first - second
	case "*":
		return first * second
	case "/":
		return first / second
	}
	return 0
}
