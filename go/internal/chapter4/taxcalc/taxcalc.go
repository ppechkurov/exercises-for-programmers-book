package taxcalc

import (
	"errors"
	"exercices/internal/input" //nolint:misspell
	"fmt"
	"io"
	"os"
)

var ErrReadOrder = errors.New("unable to read order")

type TaxCalc struct {
	in         io.Reader
	out        io.Writer
	amount     float64
	state      string
	taxPercent float64
	total      float64
}

type option func(*TaxCalc)

func WithIn(in io.Reader) option {
	return func(tc *TaxCalc) {
		tc.in = in
	}
}

func WithOut(out io.Writer) option {
	return func(tc *TaxCalc) {
		tc.out = out
	}
}

func NewTaxCalc(opts ...option) *TaxCalc {
	c := &TaxCalc{
		in:  os.Stdin,
		out: os.Stdout,
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (tc *TaxCalc) ReadInputs() (err error) {
	w := io.Discard
	isTTY := input.IsTTY(tc.in)
	if isTTY {
		w = tc.out
	}

	fmt.Fprintf(w, "What is the order amount? ")
	tc.amount, err = input.GetFloat(tc.in)
	if err != nil {
		return fmt.Errorf("invalid order amount: %w", err)
	}

	fmt.Fprintf(w, "What is the state? ")
	_, err = fmt.Fscan(tc.in, &tc.state)
	if err != nil {
		return fmt.Errorf("invalid state: %w", err)
	}

	return nil
}

const WITax = 5.5

func (tc *TaxCalc) CalculateTotal() *TaxCalc {
	if tc.state == "WI" {
		tc.taxPercent = WITax
	}

	tax := (tc.amount / 100) * tc.taxPercent
	tc.total = tc.amount + tax
	return tc
}

func (tc *TaxCalc) PrintResult() *TaxCalc {
	result := ""
	if tc.state == "WI" {
		result += fmt.Sprintf("The subtotal is $%.2f\n", tc.amount)
		result += fmt.Sprintf("The tax is $%.2f\n", tc.taxPercent)
	}
	result += fmt.Sprintf("The total is $%.2f\n", tc.total)
	fmt.Fprint(tc.out, result)
	return tc
}

func (tc *TaxCalc) Run() error {
	err := tc.ReadInputs()
	if err != nil {
		return err
	}

	tc.CalculateTotal().PrintResult()
	return nil
}

func Main() error {
	return NewTaxCalc().Run()
}
