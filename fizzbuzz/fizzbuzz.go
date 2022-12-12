package fizzbuzz

import "strconv"

type FizzBuzz struct {
}

func (b FizzBuzz) Values() []string {
	values := make([]string, 100)
	for i := 0; i < 100; i++ {
		if (i+1)%3 == 0 {
			values[i] = "Fizz"
		} else {
			values[i] = strconv.Itoa(i + 1)
		}
	}
	return values
}

func NewFizzBuzz() FizzBuzz {
	return FizzBuzz{}
}
