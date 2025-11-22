package area_test

import (
	"bytes"
	"errors"
	"exercices/area"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCalculateArea(t *testing.T) {
	t.Parallel()

	t.Run("correct inputs", func(t *testing.T) {
		t.Parallel()
		cases := map[string]struct {
			inputs string
			want   string
		}{
			"feets;length:15;width:20": {
				inputs: "feet\n15\n20",
				want: `You entered dimensions of 15 feet by 20 feet.
The area is:
300.000 square feet
27.871 square meters
`,
			},
			"meters;length:15;width:20": {
				inputs: "meter\n15\n20",
				want: `You entered dimensions of 15 meter by 20 meter.
The area is:
984.252 square feet
300.000 square meters
`,
			},
		}

		for name, tt := range cases {
			t.Run(name, func(t *testing.T) {
				buf := bytes.Buffer{}
				err := area.Main(strings.NewReader(tt.inputs), &buf)
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
		t.Parallel()
		cases := map[string]struct {
			inputs string
			want   error
		}{
			"empty length": {
				inputs: "feet\n\n15",
				want:   area.ErrEmpty,
			},
			"empty width": {
				inputs: "feet\n10\n",
				want:   area.ErrEmpty,
			},
			"invalid length": {
				inputs: "feet\nrandom_string\n15",
				want:   area.ErrInvalid,
			},
			"invalid width": {
				inputs: "feet\n10\nrandom_string",
				want:   area.ErrInvalid,
			},
			"unsupported unit": {
				inputs: "random_unit\n20\n15",
				want:   area.ErrUnsupported,
			},
		}

		for name, tt := range cases {
			t.Run(name, func(t *testing.T) {
				buf := bytes.Buffer{}
				err := area.Main(strings.NewReader(tt.inputs), &buf)
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
