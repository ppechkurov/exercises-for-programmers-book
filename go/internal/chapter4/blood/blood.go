package blood

import (
	"fmt"
	"os"
)

func Main() {
	var err error
	var code int

	defer func() {
		if err != nil {
			fmt.Println(err)
		}

		os.Exit(code)
	}()

	bc := NewCalc()
	err = bc.ReadInputs()
	if err != nil {
		code = 2
		return
	}

	BAC := bc.Calculate()
	msg := fmt.Sprintf("Your BAC is %.2f\n", BAC)
	if BAC >= 0.08 {
		msg += "It's not legal for you to drive."
	} else {
		msg += "You can drive legally."
	}
	fmt.Println(msg)
}
