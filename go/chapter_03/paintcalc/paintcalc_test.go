package paintcalc_test

import (
	"bytes"
	"exercices/paintcalc"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPaintCalc(t *testing.T) {
	t.Parallel()
	t.Run("square area", func(t *testing.T) {
		want := `You will need to purchase 2 gallons of paint to cover 360 square feet.`
		buf := bytes.Buffer{}
		err := paintcalc.Main(strings.NewReader("60\n6\n"), &buf)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		got := buf.String()
		if want != got {
			t.Error(cmp.Diff(want, got))
		}
	})
}
