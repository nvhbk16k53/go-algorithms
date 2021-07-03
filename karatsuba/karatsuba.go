package karatsuba

import (
	"bytes"
	"strings"
)

type Number struct {
	Sign   int8
	Digits []int8
}

func NewNumber(s string) Number {
	s = strings.TrimSpace(s)
	sign := int8(1)
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	num := Number{
		Sign:   sign,
		Digits: make([]int8, 0, len(s)),
	}
	for i := len(s) - 1; i >= 0; i-- {
		num.Digits = append(num.Digits, int8(s[i]-'0'))
	}
	return num
}

func (num Number) String() string {
	buf := new(bytes.Buffer)
	if num.Sign == -1 {
		buf.WriteRune('-')
	}

	n := len(num.Digits)
	for i := n - 1; i >= 0; i-- {
		buf.WriteRune(rune(num.Digits[i] + '0'))
	}

	return buf.String()
}

func int_shift(x Number, n int) Number {
	res := Number{
		Sign:   x.Sign,
		Digits: make([]int8, 0, n+len(x.Digits)),
	}

	for i := 0; i < n; i++ {
		res.Digits = append(res.Digits, 0)
	}
	res.Digits = append(res.Digits, x.Digits...)

	return res
}

func int_comp(x Number, y Number) int8 {
	if x.Sign*y.Sign < 0 {
		return x.Sign
	}

	for i := len(x.Digits) - 1; i >= 0; i-- {
		if x.Digits[i] < y.Digits[i] {
			return -x.Sign
		} else if x.Digits[i] > y.Digits[i] {
			return x.Sign
		}
	}

	return 0
}

func int_negate(x Number) Number {
	return Number{
		Sign:   -x.Sign,
		Digits: x.Digits,
	}
}

func int_add(x Number, y Number) Number {
	if len(x.Digits) > len(y.Digits) {
		ny := len(y.Digits)
		for i := 0; i < len(x.Digits)-ny; i++ {
			y.Digits = append(y.Digits, 0)
		}
	} else if len(x.Digits) < len(y.Digits) {
		nx := len(x.Digits)
		for i := 0; i < len(y.Digits)-nx; i++ {
			x.Digits = append(x.Digits, 0)
		}
	}
	sign := int8(1)
	if x.Sign*y.Sign > 0 {
		sign = x.Sign
		if sign < 0 {
			x.Sign *= -1
			y.Sign *= -1
		}
	} else {
		if x.Sign < 0 {
			if int_comp(int_negate(x), y) > 0 {
				sign = -1
				x = int_negate(x)
				y = int_negate(y)
			}
		} else {
			if int_comp(x, int_negate(y)) < 0 {
				sign = -1
				x = int_negate(x)
				y = int_negate(y)
			}
		}
	}

	res := Number{
		Sign:   sign,
		Digits: make([]int8, 0, len(x.Digits)+1),
	}
	c := int8(0)
	for i := 0; i < len(x.Digits); i++ {
		sum := x.Sign*x.Digits[i] + y.Sign*y.Digits[i] + c
		c = sum / 10
		if sum < 0 {
			c = -1
			sum = 10 + sum
		}
		res.Digits = append(res.Digits, sum%10)
	}
	if c > 0 {
		res.Digits = append(res.Digits, c)
	}
	return res
}

func int_mul(x Number, y Number) Number {
	sign := x.Sign * y.Sign
	if len(x.Digits) == 1 && len(y.Digits) == 1 {
		xy := x.Digits[0] * y.Digits[0]
		if xy == 0 {
			return Number{
				Sign:   1,
				Digits: []int8{0},
			}
		}
		if xy < 10 {
			return Number{
				Sign:   sign,
				Digits: []int8{xy},
			}
		}
		return Number{
			Sign:   sign,
			Digits: []int8{xy % 10, xy / 10},
		}
	}

	nx := len(x.Digits)
	ny := len(y.Digits)
	n := nx
	if n < ny {
		n = ny
	}
	if n%2 == 1 {
		n += 1
	}
	for i := 0; i < n-nx; i++ {
		x.Digits = append(x.Digits, 0)
	}
	for i := 0; i < n-ny; i++ {
		y.Digits = append(y.Digits, 0)
	}

	a := Number{
		Sign:   1,
		Digits: x.Digits[n/2:],
	}
	b := Number{
		Sign:   1,
		Digits: x.Digits[:n/2],
	}

	c := Number{
		Sign:   1,
		Digits: y.Digits[n/2:],
	}
	d := Number{
		Sign:   1,
		Digits: y.Digits[:n/2],
	}

	p := int_add(a, b)
	q := int_add(c, d)
	ac := int_mul(a, c)
	bd := int_mul(b, d)
	pq := int_mul(p, q)
	acbd := int_add(pq, int_negate(int_add(ac, bd)))

	return int_add(int_add(int_shift(ac, n), int_shift(acbd, n/2)), bd)
}
