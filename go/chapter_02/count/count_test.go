package count_test

import (
	"bytes"
	"exercices/count"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCount(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		input string
		want  string
	}{
		// "input Homer": {
		// 	input: "Homer",
		// 	want:  "What is the input string? Homer has 5 characters.\n",
		// },
		// "input \"\"": {
		// 	input: "",
		// 	want:  "test",
		// },
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			buf := &bytes.Buffer{}
			rc := count.Main(strings.NewReader(tt.input), buf)
			if rc != 0 {
				t.Fatalf("rc not 0: \ninput: %s", tt.input)
			}

			got := buf.String()
			if got != tt.want {
				t.Error(cmp.Diff(tt.want, got))
			}
		})
	}
}
