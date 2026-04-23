package taxcalc_test

import (
	"bytes"
	"exercices/internal/chapter4/taxcalc"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestCase[T any] map[string]T

func TestTaxCalcRun_WithValidInputs(t *testing.T) {
	t.Parallel()

	cases := TestCase[struct {
		in       []byte
		expected string
	}]{
		"amount=10 state=WI": {
			in:       []byte("10\nWI"),
			expected: "The subtotal is $10.00\nThe tax is $0.55\nThe total is $10.55\n",
		},
		"amount=10 state=MN": {
			in:       []byte("10\nMN"),
			expected: "The total is $10.00\n",
		},
		"11.067 rounded to 11.07": {
			in:       []byte("10.49\nWI"),
			expected: "The subtotal is $10.49\nThe tax is $0.58\nThe total is $11.07\n",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			in := bytes.NewReader(tt.in)
			out := bytes.NewBuffer(nil)
			err := taxcalc.NewTaxCalc(
				taxcalc.WithIn(in),
				taxcalc.WithOut(out),
			).Run()

			require.NoError(t, err)
			require.Equal(t, tt.expected, out.String())
		})
	}
}

func TestTaxCalcRun_WithInvalidInputs(t *testing.T) {
	t.Parallel()

	cases := TestCase[struct {
		in       []byte
		expected string
	}]{
		"amount=invalid": {
			in: []byte("invalid\nWI"),
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			in := bytes.NewReader(tt.in)
			err := taxcalc.NewTaxCalc(
				taxcalc.WithIn(in),
				taxcalc.WithOut(io.Discard),
			).Run()

			require.Error(t, err)
		})
	}
}
