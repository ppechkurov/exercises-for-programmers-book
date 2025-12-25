package exchange

import (
	"exercices/tty"
	"fmt"
	"io"
	"strconv"
)

func Main(in io.Reader, out io.Writer) error {
	q, rate, err := getInputs(in)
	if err != nil {
		return fmt.Errorf("getting inputs: %s", err)
	}

	res := calc(q, rate)
	fmt.Fprintf(out, "%.2f euros at an exchange rate of %.2f is\n%.2f U.S. dollars.\n", q, rate, res)
	return nil
}

func getInputs(in io.Reader) (curr, rate float64, error error) {
	input, err := tty.ReadOnce(in)
	if err != nil {
		return 0, 0, fmt.Errorf("reading quantity: %s", err)
	}
	curr, err = strconv.ParseFloat(input, 32)
	if err != nil {
		return 0, 0, err
	}

	input, err = tty.ReadOnce(in)
	if err != nil {
		return 0, 0, fmt.Errorf("reading rate: %s", err)
	}
	rate, err = strconv.ParseFloat(input, 32)
	if err != nil {
		return 0, 0, err
	}

	return curr, rate, nil
}

func calc(curr, rate float64) float64 {
	amount := curr * rate / 100
	return amount
}
