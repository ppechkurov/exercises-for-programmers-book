package madlib_test

import (
	"bytes"
	"exercices/madlib"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMadLib(t *testing.T) {
	buf := &bytes.Buffer{}
	err := madlib.Main(strings.NewReader("dog\nwalk\nblue\nquickly\n"), buf)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	want := `Enter a noun: dog
Enter a verb: walk
Enter an adjective: blue
Enter an adverb: quickly
Do you walk your blue dog quickly? That's hilarious!`

	got := buf.String()
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}
