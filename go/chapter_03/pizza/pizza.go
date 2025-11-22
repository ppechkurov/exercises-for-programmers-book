package pizza

import (
	"exercices/tty"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Main(in io.Reader, out io.Writer) error {
	if tty.IsTTY(in) {
		fmt.Fprint(out, "How many people? ")
	}

	input, err := tty.ReadOnce(in)
	if err != nil {
		return fmt.Errorf("reading people: %w", err)
	}
	people, err := strconv.Atoi(input)
	if err != nil {
		return err
	}

	if tty.IsTTY(in) {
		fmt.Fprint(out, "How many pizzas do you have? ")
	}

	input, err = tty.ReadOnce(in)
	if err != nil {
		return fmt.Errorf("reading pizzas: %w", err)
	}
	pizzas, err := strconv.Atoi(input)
	if err != nil {
		return err
	}

	fmt.Fprint(out, formatResult(people, pizzas))

	return nil
}

func formatResult(people, pizzas int) string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%d people with %d ", people, pizzas))
	if pizzas == 1 {
		b.WriteString("pizza")
	} else {
		b.WriteString("pizzas")
	}
	b.WriteString("\n")

	pieces, leftover := divide(people, pizzas)
	b.WriteString(fmt.Sprintf("Each person gets %d ", pieces))
	if pieces > 1 {
		b.WriteString("pieces")
	} else {
		b.WriteString("piece")
	}
	b.WriteString(" of pizza.\n")

	b.WriteString(fmt.Sprintf("There are %d leftover ", leftover))

	if pieces > 1 {
		b.WriteString("pieces")
	} else {
		b.WriteString("piece")
	}

	b.WriteString(".\n")
	return b.String()
}

func divide(people, pizzas int) (pieces, leftover int) {
	pieces = pizzas * 8 / people
	leftover = pizzas * 8 % people
	return pieces, leftover
}
