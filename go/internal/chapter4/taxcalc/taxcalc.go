package taxcalc

import (
	"errors"
	"exercices/internal/tty"
	"fmt"
	"io"
)

var ErrReadOrder = errors.New("unable to read order")

func Main(r io.Reader, w io.Writer) error {
	isTTY := tty.IsTTY(r)

	if isTTY {
		fmt.Fprint(w, "What's the order? ")
	}

	order := 0
	leftover := ""
	_, err := fmt.Fscanf(r, "%d", &order)
	if err != nil {
		fmt.Fscanf(r, "%s", &leftover)
		return fmt.Errorf("%w: read %q: %w", ErrReadOrder, leftover, err)
	}

	if isTTY {
		fmt.Fprintln(w, order)
	}
	return nil
}
