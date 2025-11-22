package retire_test

import (
	"bytes"
	"exercices/retire"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRetire(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	input := "25\n65"
	want := `What is your current age? 25
At what age would you like to retire? 65
You have 40 years left until you can retire.
It's 2025, so you can retire in 2065.`

	err := retire.Main(strings.NewReader(input), buf)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	got := buf.String()
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}

func TestRetireNegative(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	input := "65\n25"
	want := `What is your current age? 65
At what age would you like to retire? 25
You should've been retired 40 years ago! What are you waiting for?`

	err := retire.Main(strings.NewReader(input), buf)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	got := buf.String()
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}
