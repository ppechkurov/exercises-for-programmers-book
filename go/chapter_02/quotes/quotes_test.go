package quotes_test

import (
	"bytes"
	"exercices/quotes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQuotes(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		input string
		want  string
	}{
		"valid inputs": {
			input: `These aren't the droids you're looking for.
Obi-Wan Kenobi
`,
			want: `What is the quote? These aren't the droids you're looking for.
Who said it? Obi-Wan Kenobi
Obi-Wan Kenobi says, "These aren't the droids you're looking for."
`,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			buf := &bytes.Buffer{}
			err := quotes.Main(strings.NewReader(tt.input), buf)
			if err != nil {
				t.Logf("err: %s", err)
				t.Fatalf("unexpected error: \ninput: %s", tt.input)
			}

			got := buf.String()
			if got != tt.want {
				t.Error(cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestChalleng(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		input []quotes.Quote
		want  string
	}{
		"valid inputs": {
			input: []quotes.Quote{
				{
					Author: "Obi-Wan Kenobi",
					Quote:  "These aren't the droids you're looking for.",
				},
			},
			want: "Obi-Wan Kenobi says, \"These aren't the droids you're looking for.\"\n",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			buf := &bytes.Buffer{}
			quotes.Challenge(tt.input, buf)

			got := buf.String()
			if got != tt.want {
				t.Error(cmp.Diff(tt.want, got))
			}
		})
	}
}
