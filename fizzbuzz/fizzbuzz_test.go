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
