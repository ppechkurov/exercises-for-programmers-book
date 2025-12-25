package sinterest

import (
	"exercices/input"
	"exercices/tty"
	"fmt"
	"io"
	"math"
)

func Main(in io.Reader, out io.Writer) error {
	if tty.IsTTY(in) {
		fmt.Fprint(out, "Enter the principal: ")
	}
	principal, err := input.GetInt(in)
	if err != nil {
		return fmt.Errorf("reading principal: %w", err)
	}

	if tty.IsTTY(in) {
		fmt.Fprint(out, "Enter the rate of interest: ")
	}
	rate, err := input.GetFloat(in)
	if err != nil {
		return fmt.Errorf("reading rate: %w", err)
	}

	if tty.IsTTY(in) {
		fmt.Fprint(out, "Enter the number of years: ")
	}
	years, err := input.GetInt(in)
	if err != nil {
		return fmt.Errorf("reading years: %w", err)
	}

	for i := range years + 1 {
		if i == 0 {
			continue
		}
		result := CalculateInterest(principal, i, rate)
		fmt.Fprint(out, SPrintResult(i, rate, result))
	}

	return nil
}

func CalculateInterest(principal, years int, rate float64) float64 {
	interest := float64(principal) * (1 + rate/100*float64(years))
	return math.Round(interest*100) / 100
}

func SPrintResult(years int, rate, result float64) string {
	return fmt.Sprintf("After %v years at %v%%, the investment will be worth $%v.\n", years, rate, result)
}
