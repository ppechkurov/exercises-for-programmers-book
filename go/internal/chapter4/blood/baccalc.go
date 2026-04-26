package blood

import (
	"errors"
	"exercices/internal/input"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

type BACCalculator struct {
	weight float64
	gender Gender
	drinks float64
	oz     float64
	hours  float64

	in  io.Reader
	out io.Writer
}

type option func() BACCalculator

func NewCalc(opts ...option) *BACCalculator {
	bc := &BACCalculator{
		in:  os.Stdin,
		out: os.Stdout,
	}

	for _, opt := range opts {
		opt()
	}

	return bc
}

func (bc *BACCalculator) ReadInputs() error {
	err := bc.readWeight()
	if err != nil {
		if err == io.EOF {
			return errors.New("no inputs provided")
		}

		return err
	}

	err = bc.readGender()
	if err != nil {
		return err
	}

	err = bc.readNumberOfDrinks()
	if err != nil {
		return err
	}

	err = bc.readAlcoholOz()
	if err != nil {
		return err
	}

	err = bc.readHours()
	if err != nil {
		return err
	}

	return nil
}

func (bc *BACCalculator) Calculate() (bac float64) {
	r := 0.73
	if bc.gender == "female" {
		r = 0.66
	}

	return (bc.oz * bc.drinks * 5.14 / bc.weight * r) - 0.15*bc.hours
}

func (bc *BACCalculator) readWeight() error {
	bc.prompt("What is your weight? ")
	var err error
	bc.weight, err = input.GetFloat(bc.in)
	return err
}

func (bc *BACCalculator) readGender() error {
	bc.prompt("What is your gender? ")
	_, err := fmt.Fscan(bc.in, &bc.gender)
	if err != nil {
		return err
	}

	switch bc.gender {
	case GenderFemale, GenderMale:
		return nil
	default:
		return fmt.Errorf(
			"invalid gender: either %q or %q expected",
			GenderMale, GenderFemale,
		)
	}
}

func (bc *BACCalculator) readNumberOfDrinks() (err error) {
	bc.prompt("What is the number of drinks? ")
	bc.drinks, err = input.GetFloat(bc.in)
	if errors.Is(err, strconv.ErrSyntax) {
		return errors.New("invalid number of drinks")
	}
	return err
}

func (bc *BACCalculator) readAlcoholOz() (err error) {
	bc.prompt("What is the amount of alcohol (oz) in each drink? ")
	bc.oz, err = input.GetFloat(bc.in)
	if errors.Is(err, strconv.ErrSyntax) {
		return errors.New("invalid alcohol amount")
	}
	return err
}

func (bc *BACCalculator) readHours() (err error) {
	bc.prompt("How many hours ago since last drink? ")
	bc.hours, err = input.GetFloat(bc.in)
	if errors.Is(err, strconv.ErrSyntax) {
		return errors.New("invalid hours")
	}
	return err
}

func (bc *BACCalculator) prompt(msg string) {
	if input.IsTTY(bc.in) {
		fmt.Fprint(bc.out, msg)
	}
}
