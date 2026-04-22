package taxcalc

import (
	"errors"
	"exercices/internal/input" //nolint:misspell
	"fmt"
	"io"
)

var ErrReadOrder = errors.New("unable to read order")

func Main(in io.Reader, out io.Writer) error {
	w := io.Discard
	isTTY := input.IsTTY(in)
	if isTTY {
		w = out
	}

	amount, state, err := GetInputs(in, w)
	if err != nil {
		return err
	}

	tax := CalculateTax(state)
	total := CalculateTotal(amount, tax)

	result := ""
	if state == "WI" {
		result += fmt.Sprintf("The subtotal is $%.2f\n", amount)
		result += fmt.Sprintf("The tax is $%.2f\n", tax)
	}
	result += fmt.Sprintf("The total is $%.2f\n", total)
	fmt.Fprint(out, result)

	return nil
}

const (
	WITax = 5.5
	MNTax = 0.0
)

func GetInputs(r io.Reader, w io.Writer) (
	order float64,
	state string,
	err error,
) {
	fmt.Fprintf(w, "What is the order amount? ")
	order, err = input.GetFloat(r)
	if err != nil {
		return 0, "", fmt.Errorf("invalid order amount: %w", err)
	}

	state = ""
	fmt.Fprintf(w, "What is the state? ")
	_, err = fmt.Fscan(r, &state)
	if err != nil {
		return 0, "", err
	}

	return order, state, err
}

func CalculateTax(state string) float64 {
	if state == "WI" {
		return 5.5
	}
	return 0.0
}

func CalculateTotal(amount float64, taxPercent float64) float64 {
	tax := (amount / 100) * taxPercent
	return amount + tax
}
