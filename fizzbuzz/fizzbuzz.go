package fizzbuzz

import "strconv"

type FizzBuzz struct {
}

func (b FizzBuzz) Values() []string {
	values := make([]string, 100)
	for i := 0; i < 100; i++ {
		values[i] = strconv.Itoa(i + 1)
	}
	return values
}

func NewFizzBuzz() FizzBuzz {
	return FizzBuzz{}
}
