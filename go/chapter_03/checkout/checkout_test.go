package checkout_test

import (
	"bytes"
	"exercices/checkout"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCheckout(t *testing.T) {
	t.Parallel()
	t.Run("default checkout", func(t *testing.T) {
		t.Parallel()
		buf := &bytes.Buffer{}

		inputs := []string{
			"25", // price of item 1
			"2",  // quantity of item 1
			"10", // price of item 2
			"1",  // quantity of item 2
			"4",  // price of item 3
			"1",  // quantity of item 3
			"5",  // price of item 4
			"2",  // quantity of item 4
		}
		want := `Subtotal: $74.00
Tax: $4.07
Total: $78.07
`
		err := checkout.Main(strings.NewReader(strings.Join(inputs, "\n")), buf)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		got := buf.String()
		if want != got {
			t.Error(cmp.Diff(want, got))
		}
	})
}
