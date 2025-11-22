package paintcalc

import (
	room "exercices/internal"
	"exercices/tty"
	"fmt"
	"io"
	"strconv"
)

func Main(in io.Reader, out io.Writer) error {
	params, err := getParams(in, out)
	if err != nil {
		return fmt.Errorf("getting params: %w", err)
	}

	r := room.NewRoom(params[0], params[1])
	fmt.Fprint(out, r)
	return nil
}

func getParams(r io.Reader, w io.Writer) ([]int, error) {
	inputs := make([]int, 2)
	for i, v := range []string{"length", "width"} {
		if tty.IsTTY(r) {
			fmt.Fprintf(w, "%s: ", v)
		}
		input, err := tty.ReadOnce(r)
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", v, err)
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			return nil, fmt.Errorf("parsing %s: %w", v, err)
		}
		inputs[i] = num
	}
	return inputs, nil
}
