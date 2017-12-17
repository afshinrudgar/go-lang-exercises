package clock

import "fmt"

type Clock struct {
	Hour   int
	Minute int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

func newByMinutes(totalMinutes int) Clock {
	for totalMinutes < 0 {
		totalMinutes += (24 * 60)
	}
	totalMinutes %= (24 * 60)
	return Clock{totalMinutes / 60, totalMinutes % 60}
}

func New(hour, minute int) Clock {
	return newByMinutes(hour*60 + minute)
}

func (c Clock) Add(minutes int) Clock {
	totalMinutes := c.Hour*60 + c.Minute
	totalMinutes += minutes
	return newByMinutes(totalMinutes)
}
