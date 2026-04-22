package input

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var ErrEmpty = errors.New("empty input")

func IsTTY(r io.Reader) bool {
	f, ok := r.(*os.File)
	if !ok {
		return false
	}
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

func GetInt(in io.Reader) (int, error) {
	input, err := readOnce(in)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(input)
}

func GetFloat(in io.Reader) (float64, error) {
	input, err := readOnce(in)
	if err != nil {
		return 0, err
	}

	res, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to parse %q: wanted float", input)
	}

	return res, nil
}

func readOnce(r io.Reader) (res string, err error) {
	_, err = fmt.Fscan(r, &res)
	if err != nil {
		return "", err
	}
	if res == "" {
		err = ErrEmpty
	}
	return res, err
}
