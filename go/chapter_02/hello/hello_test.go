package hello_test

import (
	"bytes"
	"exercices/hello"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHelloWithInvalidInputs(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in   string
		out  string
		want int
	}{
		"empty input": {
			in:   "",
			want: 1,
			out:  "read empty string from input",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			out := &bytes.Buffer{}
			got := hello.Main(bytes.NewReader([]byte(tt.in)), out)
			if tt.want != got {
				t.Error(cmp.Diff(tt.want, got))
				t.Error(cmp.Diff(tt.out, out.String()))
			}
		})
	}
}

func TestHelloWithValidInputs(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in   string
		want string
		rc   int
	}{
		"returns correct greetings": {
			in:   "Test",
			want: "What is your name? Hello, Test, nice to meet you!\n",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			buf := &bytes.Buffer{}
			rc := hello.Main(strings.NewReader(tt.in), buf)
			if rc != tt.rc {
				t.Fatalf("rc not 0: \nin: %s", tt.in)
			}

			got := buf.String()
			if got != tt.want {
				t.Error(cmp.Diff(tt.want, got))
			}
		})
	}
}
