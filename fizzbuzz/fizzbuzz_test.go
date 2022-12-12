package fizzbuzz_test

import (
	"fizzbuzz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Returns_100_Values(t *testing.T) {
	fizzBuzz := fizzbuzz.NewFizzBuzz()

	values := fizzBuzz.Values()

	assert.Equal(t, 100, len(values))
}

func Test_Numbers_not_multiple_does_not_change(t *testing.T) {
	fizzBuzz := fizzbuzz.NewFizzBuzz()

	values := fizzBuzz.Values()

	assert.Equal(t, "1", values[0])
	assert.Equal(t, "2", values[1])
	assert.Equal(t, "4", values[3])
}
