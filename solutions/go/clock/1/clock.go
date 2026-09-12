package clock

import (
    "fmt"
)
// Define the Clock type here.
type Clock struct {
    hour int
    min int
}

func New(h, m int) Clock {
	total :=(m+h*60)%1440
    if total < 0 {
        total += 1440
    }
    return Clock{
        hour: total/60,
        min: total % 60,
    }
}

func (c Clock) Add(m int) Clock {
	m += c.hour*60 + c.min
    return New(0, m)
}

func (c Clock) Subtract(m int) Clock {
	total := c.hour*60 + c.min - m
    return New(0, total)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}
