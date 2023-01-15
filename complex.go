package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

type complex struct {
	real, imaginary *big.Float
}

type complexFloat struct {
	real, imaginary float64
}

func newComplex(r, i float64) *complex {
	return &complex{big.NewFloat(r), big.NewFloat(i)}
}

func randomComplex(fr, tr, fi, ti float64) *complex {
	return newComplex(
		fr+rand.Float64()*(tr-fr),
		fi+rand.Float64()*(ti-fi),
	)
}

func (c *complex) setEqualTo(c2 *complex) {
	c.real.Copy(c2.real)
	c.imaginary.Copy(c2.imaginary)
}

func multiply(a, b *complex) *complex {
	// (a + bi) * (c + di) = ac + bci + adi + bdii = ac + bci + adi - bd = (ac - bd) + (bc+ad)i
	result := newComplex(0, 0)
	left := newComplex(0, 0)
	right := newComplex(0, 0)
	left.real.Mul(a.real, b.real)
	right.real.Mul(a.imaginary, b.imaginary)
	left.imaginary.Mul(a.imaginary, b.real)
	right.imaginary.Mul(a.real, b.imaginary)
	result.real.Sub(left.real, right.real)
	result.imaginary.Add(left.imaginary, right.imaginary)
	//return &complex{
	//	real:      a.real*b.real - a.imaginary*b.imaginary,
	//	imaginary: a.imaginary*b.real + a.real*b.imaginary,
	//}
	return result
}

func power(a *complex, p int) *complex {
	result := newComplex(0, 0)
	result.setEqualTo(a)
	for i := 0; i < p-1; i++ {
		result = multiply(result, a)
	}
	return result
}

func sum(a, b *complex) *complex {
	result := newComplex(0, 0)
	result.real.Add(a.real, b.real)
	result.imaginary.Add(a.imaginary, b.imaginary)
	return result
}

func sub(a, b *complex) *complex {
	result := newComplex(0, 0)
	result.real.Sub(a.real, b.real)
	result.imaginary.Sub(a.imaginary, b.imaginary)
	return result
}

func addFloat(a *complex, bReal, bIm float64) *complex {
	return sum(a, newComplex(bReal, bIm))
}

func subFloat(a *complex, bReal, bIm float64) *complex {
	return sub(a, newComplex(bReal, bIm))
}

func mulFloat(a *complex, bReal, bIm float64) *complex {
	return multiply(a, newComplex(bReal, bIm))
}

func (c *complex) toString() string {
	return fmt.Sprintf("(%.4f + %.4fi)", c.real, c.imaginary)
}

func (c *complex) squareMagnitude() float64 {
	result := newComplex(0, 0)
	result.real.Mul(c.real, c.real)
	result.imaginary.Mul(c.imaginary, c.imaginary)
	ret := big.NewFloat(0)
	ret.Add(result.real, result.imaginary)
	floatRes, _ := ret.Float64()
	return floatRes // c.real*c.real + c.imaginary*c.imaginary
}
