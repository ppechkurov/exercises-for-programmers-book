package input

import (
	"exercices/tty"
	"io"
	"strconv"
)

func GetInt(in io.Reader) (int, error) {
	input, err := tty.ReadOnce(in)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(input)
}

func GetFloat(in io.Reader) (float64, error) {
	input, err := tty.ReadOnce(in)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(input, 64)
}
