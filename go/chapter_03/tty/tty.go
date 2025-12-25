package tty

import (
	"errors"
	"fmt"
	"io"
	"os"

	"golang.org/x/term"
)

var ErrEmpty = errors.New("empty input")

func IsTTY(r io.Reader) bool {
	f, ok := r.(*os.File)
	return ok && term.IsTerminal(int(f.Fd()))
}

func ReadOnce(r io.Reader) (res string, err error) {
	_, err = fmt.Fscan(r, &res)
	if err != nil {
		return "", err
	}
	if res == "" {
		err = ErrEmpty
	}
	return res, err
}
