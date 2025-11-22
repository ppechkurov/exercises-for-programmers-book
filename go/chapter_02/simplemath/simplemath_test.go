package simplemath_test

import (
	"bytes"
	"errors"
	"exercices/simplemath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSimpleMath(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		want      string
		inputs    string
		wantError error
	}{
		"works with valid inputs": {
			want: `What is the first number? 10
What is the second number? 5
10 + 5 = 15
10 - 5 = 5
10 * 5 = 50
10 / 5 = 2
`,
			inputs: "10\n5",
		},
		"catches an empty input": {
			want:      "What is the first number? ",
			wantError: simplemath.ErrEmpty,
		},
		"catches invalid input": {
			inputs:    "test\nanother test",
			want:      "What is the first number? ",
			wantError: simplemath.ErrInvalid,
		},
		"catches negative input": {
			inputs:    "-1\n-10",
			want:      "What is the first number? ",
			wantError: simplemath.ErrNegative,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			buf := &bytes.Buffer{}
			err := simplemath.Main(strings.NewReader(tt.inputs), buf)
			if err != nil {
				if tt.wantError == nil {
					t.Fatalf("unexpected error: %s", err)
				}
				if !errors.Is(err, tt.wantError) {
					t.Fatalf("wanted error: %s, but got: %s", tt.wantError, err)
				}
			}

			got := buf.String()
			if tt.want != got {
				t.Error(cmp.Diff(tt.want, got))
			}
		})
	}
}
