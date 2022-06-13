package main

import (
	"fmt"
	"math/rand"
)

type complex struct {
	real, imaginary float64
}

func newComplex(r, i float64) *complex {
	return &complex{r, i}
}

func randomComplex(fr, tr, fi, ti float64) *complex {
	return &complex{
		real:      fr + rand.Float64() * (tr - fr),
		imaginary: fi + rand.Float64() * (ti - fi),
	}
}

func (c *complex) setEqualTo(c2 *complex) {
	c.real = c2.real
	c.imaginary = c2.imaginary
}

func multiply(a, b *complex) *complex {

	// (a + bi) * (c + di) = ac + bci + adi + bdii = ac + bci + adi - bd = (ac - bd) + (bc+ad)i
	return &complex{
		real:      a.real*b.real - a.imaginary*b.imaginary,
		imaginary: a.imaginary*b.real + a.real*b.imaginary,
	}
}

func sum(a, b *complex) *complex {
	return &complex{
		real:      a.real+b.real,
		imaginary: a.imaginary+b.imaginary,
	}
}

func (c *complex) toString() string {
	return fmt.Sprintf("(%.4f + %.4fi)", c.real, c.imaginary)
}

func (c *complex) squareMagnitude() float64 {
	return c.real * c.real + c.imaginary * c.imaginary
}
