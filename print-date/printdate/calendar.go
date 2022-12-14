package printdate

import "time"

type Calendar interface {
	Today() string
}

func NewCalendar() Calendar {
	return calendar{}
}

type calendar struct {
}

func (calendar calendar) Today() string {
	return time.Now().String()
}
