package room

import "fmt"

const Cover = 350

type Room struct {
	length, width int
}

func NewRoom(length, width int) Room {
	return Room{length, width}
}

func (r Room) Area() int {
	return r.length * r.width
}

func (r Room) Gallons() int {
	area := r.Area()
	gal := area / Cover
	mod := area % Cover
	if mod > 0 {
		gal += 1
	}
	return gal
}

func (r Room) String() string {
	return fmt.Sprintf(
		"You will need to purchase %d gallons of paint to cover %d square feet.",
		r.Gallons(), r.Area())
}
