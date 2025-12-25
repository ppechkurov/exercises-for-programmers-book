package sinterest_test

import (
	"bytes"
	"exercices/sinterest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSinterest(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		input []string
		want  string
	}{
		"default": {
			input: []string{"1500", "4.3", "4"},
			want: `After 1 years at 4.3%, the investment will be worth $1564.5.
After 2 years at 4.3%, the investment will be worth $1629.
After 3 years at 4.3%, the investment will be worth $1693.5.
After 4 years at 4.3%, the investment will be worth $1758.
`,
		},
		"round next penny": {
			input: []string{"1503", "4.3", "4"},
			want: `After 1 years at 4.3%, the investment will be worth $1567.63.
After 2 years at 4.3%, the investment will be worth $1632.26.
After 3 years at 4.3%, the investment will be worth $1696.89.
After 4 years at 4.3%, the investment will be worth $1761.52.
`,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			buf := bytes.Buffer{}
			input := strings.Join(tt.input, " ")
			err := sinterest.Main(strings.NewReader(input), &buf)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got := buf.String()
			if tt.want != got {
				t.Error(cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestPrintResult(t *testing.T) {
	t.Parallel()

	years := 4
	rate := 4.30
	result := float64(1758)

	want := "After 4 years at 4.3%, the investment will be worth $1758.\n"
	got := sinterest.SPrintResult(years, rate, result)

	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCalculateInterest(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		principal int
		years     int
		rate      float64
		want      float64
	}{
		"default": {
			principal: 1500,
			years:     4,
			rate:      4.3,
			want:      1758,
		},
		"rounded to next penny": {
			principal: 1500,
			years:     4,
			rate:      4.8888,
			want:      1793.33,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			want := tt.want
			got := sinterest.CalculateInterest(
				tt.principal, tt.years, tt.rate)

			if want != got {
				t.Error(cmp.Diff(want, got))
			}
		})
	}
}
