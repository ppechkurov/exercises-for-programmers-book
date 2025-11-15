package parser

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Parser struct {
	input io.Reader
}

func New() *Parser {
	return &Parser{
		input: os.Stdin,
	}
}

func (p *Parser) WithInput(in io.Reader) *Parser {
	p.input = in
	return p
}

func (p *Parser) Parse() (float32, error) {
	read, err := readInput(p.input)
	if err != nil {
		return 0.0, err
	}

	result, err := strconv.ParseFloat(read, 32)
	if err != nil {
		return 0.0, fmt.Errorf("unable to parse: %w", err)
	}

	return float32(result), nil
}

func readInput(in io.Reader) (string, error) {
	s := bufio.NewScanner(in)
	if !s.Scan() || s.Text() == "" {
		return "", errors.New("unable to read input")
	}
	return s.Text(), nil
}
