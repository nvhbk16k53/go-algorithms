package karatsuba

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKaratsuba(t *testing.T) {
	for i := 0; i < 100; i++ {
		x := int(rand.Int31() + 1)
		y := int(rand.Int31() + 1)
		expect := x * y

		num1 := NewNumber(strconv.Itoa(x))
		num2 := NewNumber(strconv.Itoa(y))
		p := int_mul(num1, num2)
		p_int, err := strconv.Atoi(p.String())

		assert.NoError(t, err)
		assert.Equal(t, expect, p_int, "Calculate %d * %d, expect %d, got %d", x, y, expect, p_int)
	}
}
