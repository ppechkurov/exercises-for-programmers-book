package area

import (
	"errors"
	"exercices/tty"
	"fmt"
	"io"
	"strconv"
)

var (
	ErrEmpty       = errors.New("empty input")
	ErrInvalid     = errors.New("input invalid")
	ErrUnsupported = errors.New("unsupported unit")
)

const (
	sqareConv   = 0.09290304
	meterToFeet = 3.28084
)

type Inputs struct {
	Length string
	Width  string
}

func Main(in io.Reader, out io.Writer) error {
	if tty.IsTTY(in) {
		fmt.Fprintf(out, "What is the unit? ")
	}
	unit, err := readOnce(in)
	if err != nil {
		return fmt.Errorf("reading %s: %w", "unit", err)
	}

	if unit == "" {
		unit = "meter"
	}
	if unit != "meter" && unit != "feet" {
		return fmt.Errorf("%q: %w", unit, ErrUnsupported)
	}

	inputs := make([]int, 2)
	for i, name := range []string{"length", "width"} {
		if tty.IsTTY(in) {
			fmt.Fprintf(out, "What is the %s of the room in feet? ", name)
		}

		input, err := readOnce(in)
		if err != nil {
			return fmt.Errorf("reading %s: %w", name, err)
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			return fmt.Errorf("parsing %s %q: %w", name, input, ErrInvalid)
		}

		inputs[i] = num
	}

	length, width := inputs[0], inputs[1]
	fmt.Fprintf(out,
		"You entered dimensions of %d %s by %d %s.\n",
		length, unit, width, unit)

	var meters float32
	var feets float32
	if unit == "meter" {
		meters = float32(length) * float32(width)
		feets = meters * meterToFeet
	} else {
		feets = float32(length) * float32(width)
		meters = feets * sqareConv
	}

	fmt.Fprintf(out, "The area is:\n")
	fmt.Fprintf(out, "%.3f square feet\n", feets)
	fmt.Fprintf(out, "%.3f square meters\n", meters)

	return nil
}

func readOnce(r io.Reader) (res string, err error) {
	_, err = fmt.Fscan(r, &res)
	if err != nil && err != io.EOF {
		return "", err
	}
	if res == "" {
		err = ErrEmpty
	}
	return res, err
}

func getFeet(length, width float32) float32 {
	return length * width
}

func getMeters(feet float32) float32 {
	return sqareConv * feet
}
