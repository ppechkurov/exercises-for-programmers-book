package taxcalc_test

import (
	"bytes"
	"exercices/internal/chapter4/taxcalc"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	cases := TestCase[struct {
		in       []byte
		expected string
	}]{
		"amount=10 state=WI": {
			in:       []byte("10\nWI"),
			expected: "The subtotal is $10.00\nThe tax is $5.50\nThe total is $10.55\n",
		},
		"amount=10 state=MN": {
			in:       []byte("10\nMN"),
			expected: "The total is $10.00\n",
		},
		"11.067 rounded to 11.07": {
			in:       []byte("10.49\nWI"),
			expected: "The subtotal is $10.49\nThe tax is $5.50\nThe total is $11.07\n",
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			out := bytes.NewBuffer(nil)
			err := taxcalc.Main(bytes.NewReader(tt.in), out)
			require.NoError(t, err)
			require.Equal(t, tt.expected, out.String())
		})
	}
}

func TestCalculateTotal(t *testing.T) {
	tests := []struct {
		name       string
		amount     float64
		taxPercent float64
		want       float64
	}{
		{
			name:       "tax=5.5%",
			amount:     10,
			taxPercent: 5.5,
			want:       10.55,
		},
		{
			name:       "tax=0.0%",
			amount:     10,
			taxPercent: 0,
			want:       10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := taxcalc.CalculateTotal(tt.amount, tt.taxPercent)
			assert.Equal(t, tt.want, got)
		})
	}
}

type TestCase[T any] map[string]T

func TestValidGetInputs(t *testing.T) {
	cases := TestCase[struct {
		in    []byte
		order float64
		state string
	}]{
		"10 WI": {
			in:    []byte("10\nWI"),
			order: 10,
			state: "WI",
		},
		"10.0 MN": {
			in:    []byte("10.0\nMN"),
			order: 10.0,
			state: "MN",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			w := bytes.NewBuffer(nil)
			order, state, err := taxcalc.GetInputs(bytes.NewReader(tt.in), w)
			require.NoError(t, err)
			assert.Equal(t, tt.order, order)
			assert.Equal(t, tt.state, state)
		})
	}
}

func TestNOTValidGetInputs(t *testing.T) {
	cases := TestCase[struct {
		in []byte
	}]{
		"invalid order returns error": {
			in: []byte("invalid order"),
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			w := bytes.NewBuffer(nil)
			_, _, err := taxcalc.GetInputs(bytes.NewReader(tt.in), w)
			require.Error(t, err)
		})
	}
}

func TestCalculateTax(t *testing.T) {
	cases := TestCase[struct {
		state    string
		expected float64
	}]{
		"returns 5.5 for WI":   {state: "WI", expected: 5.5},
		"returns 0 for others": {state: "MN", expected: 0.0},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := taxcalc.CalculateTax(tt.state)
			require.Equal(t, tt.expected, actual)
		})
	}
}
