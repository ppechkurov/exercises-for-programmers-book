package pizza_test

import (
	"bytes"
	"errors"
	"exercices/pizza"
	"exercices/tty"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestPizzas(t *testing.T) {
	t.Parallel()

	t.Run("correct inputs", func(t *testing.T) {
		t.Parallel()
		cases := map[string]struct {
			inputs string
			want   string
		}{
			"people:8;pizzas:2": {
				inputs: "8\n2",
				want: `8 people with 2 pizzas
Each person gets 2 pieces of pizza.
There are 0 leftover pieces.
`,
			},
			"people:7;pizzas:1": {
				inputs: "7\n1",
				want: `7 people with 1 pizza
Each person gets 1 piece of pizza.
There are 1 leftover piece.
`,
			},
		}

		for name, tt := range cases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				buf := bytes.Buffer{}
				err := pizza.Main(strings.NewReader(tt.inputs), &buf)
				if err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				got := buf.String()
				if tt.want != got {
					t.Error(cmp.Diff(tt.want, got))
				}
			})
		}
	})

	t.Run("invalid inputs", func(t *testing.T) {
		cases := map[string]struct {
			inputs string
			want   error
		}{
			"empty people": {
				inputs: "\n15",
				want:   tty.ErrEmpty,
			},
			"empty pizzas": {
				inputs: "10\n",
				want:   tty.ErrEmpty,
			},
		}

		for name, tt := range cases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				buf := bytes.Buffer{}
				err := pizza.Main(strings.NewReader(tt.inputs), &buf)
				if nil == err {
					t.Fatalf("wanted error but didn't get one: %s", tt.inputs)
				}

				if !errors.Is(err, tt.want) {
					t.Error(cmp.Diff(tt.want, err, cmpopts.EquateErrors()))
				}
			})
		}
	})
}
