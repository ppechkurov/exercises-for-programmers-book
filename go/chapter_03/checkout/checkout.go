package checkout

import (
	"errors"
	"exercices/tty"
	"fmt"
	"io"
	"strconv"
)

type Input struct {
	price    float64
	quantity int
}

func Main(in io.Reader, out io.Writer) error {
	inputs, err := getInputs(in)
	if err != nil {
		return err
	}

	s := calcSubtotal(inputs)
	tax := s * 5.5 / 100
	total := s + tax

	fmt.Fprintf(out,
		`Subtotal: $%.2f
Tax: $%.2f
Total: $%.2f
`, s, tax, total)
	return nil
}

func getInputs(r io.Reader) ([]Input, error) {
	inputs := []Input{}

	i := 1
	for {
		s, err := tty.ReadOnce(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("reading price of item %d: %s", i, err)
		}

		priceResult, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return nil, err
		}

		s, err = tty.ReadOnce(r)
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("reading quantity of item %d: %s", i, err)
		}
		quantityResult, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		inputs = append(inputs, Input{
			price:    priceResult,
			quantity: quantityResult,
		})

		i++
	}
	return inputs, nil
}

func calcSubtotal(inputs []Input) float64 {
	s := 0.0
	for _, v := range inputs {
		s += v.price * float64(v.quantity)
	}
	return s
}
