package retire

import (
	"errors"
	"exercices/tty"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"time"
)

func Main(in io.Reader, out io.Writer) error {
	promts := []string{
		"What is your current age? ",
		"At what age would you like to retire? ",
	}

	answers := make([]int, len(promts))
	for i, p := range promts {
		fmt.Fprint(out, p)
		var ans string
		_, err := fmt.Fscan(in, &ans)
		if err != nil && err != io.EOF {
			return err
		}
		if ans == "" {
			return errors.New("input empty")
		}
		if !tty.IsTTY() {
			fmt.Fprint(out, ans+"\n")
		}

		num, err := strconv.Atoi(ans)
		if err != nil {
			return err
		}
		answers[i] = num
	}

	current := time.Now().Year()
	log.Printf("current: %v\n", current)
	left := answers[1] - answers[0]
	if left < 0 {
		fmt.Fprintf(out, "You should've been retired %d years ago! What are you waiting for?", int(math.Abs(float64(left))))
		return nil
	}

	fmt.Fprintf(out, "You have %d years left until you can retire.\n", left)
	fmt.Fprintf(out, "It's %d, so you can retire in %d.", current, current+left)

	return nil
}
