package exchange_test

import (
	"bytes"
	"exercices/exchange"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExchange(t *testing.T) {
	t.Parallel()

	t.Run("default case", func(t *testing.T) {
		t.Parallel()

		buf := &bytes.Buffer{}
		inputs := "81\n137.51\n"
		err := exchange.Main(strings.NewReader(inputs), buf)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		want := `81.00 euros at an exchange rate of 137.51 is
111.38 U.S. dollars.
`

		got := buf.String()
		if want != got {
			t.Error(cmp.Diff(want, got))
		}
	})
}

// func TestCalc(t *testing.T) {
// 	want := 111.38
// 	got := exchange.Calc(81, 137.51)
// 	if want != got {
// 		t.Error(cmp.Diff(want, got))
// 	}
// }
