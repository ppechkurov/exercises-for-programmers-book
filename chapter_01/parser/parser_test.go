package parser_test

import (
	"exercices/parser"
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name      string
		input     io.Reader
		want      float32
		wantError bool
	}{
		{
			name:  "parses number",
			input: strings.NewReader("10.5"),
			want:  10.5,
		},
		{
			name:      "returns error when invalid input",
			input:     strings.NewReader("invalid input"),
			wantError: true,
		},
		{
			name:      "returns error when got an empty string",
			input:     strings.NewReader("\n"),
			wantError: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := parser.New().WithInput(tt.input).Parse()
			if err != nil {
				if !tt.wantError {
					t.Errorf("unexpected error: %s", err)
				}
				return
			}

			if tt.want != got {
				t.Error(cmp.Diff(tt.want, got))
			}
		})
	}
}
