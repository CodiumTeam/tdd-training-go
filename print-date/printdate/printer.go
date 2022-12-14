package printdate

import "fmt"

type Printer interface {
	PrintLine(line string)
}

func NewPrinter() Printer {
	return printer{}
}

type printer struct {
}

func (printer printer) PrintLine(line string) {
	fmt.Println(line)
}
