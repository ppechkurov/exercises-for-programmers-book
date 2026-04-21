package taxcalc

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var ErrReadOrder = errors.New("unable to read order")

func Main(r io.Reader, w io.Writer) error {
	fmt.Fprint(w, "What's the order? ")

	order := 0
	leftover := ""
	_, err := fmt.Fscanf(r, "%d", &order)
	if err != nil {
		fmt.Fscanf(r, "%s", &leftover)
		return fmt.Errorf("%w: read %q: %w", ErrReadOrder, leftover, err)
	}

	file, ok := r.(*os.File)
	fi, _ := file.Stat()
	if ok && fi.Mode()&os.ModeCharDevice != 0 {
		fmt.Fprint(w, order)
	}
	fmt.Fprint(w, order)
	return nil
}
