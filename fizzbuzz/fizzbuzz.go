package fizzbuzz

type FizzBuzz struct {
}

func (b FizzBuzz) Values() []string {
	return make([]string, 100)
}

func NewFizzBuzz() FizzBuzz {
	return FizzBuzz{}
}
