package main

import (
	"math"
	"math/rand"
)

type fractalType uint8

const (
	MANDELBROT fractalType = iota
	JULIA_MB
	TRICORN
	JULIA_TRICORN
	BURNING_SHIP
	JULIA_BS
	TOTAL_FRACTALS
)

var maxSetCheckIterations = 20
var orderOfFractalExpression = 2

func (c1 *complex) iterateAsZ(c *complex) {
	switch currentFractal {
	//case MANDELBROT, JULIA_MB:
	//	break
	case TRICORN, JULIA_TRICORN:
		c1.imaginary = -c1.imaginary
	case BURNING_SHIP, JULIA_BS:
		c1.real = math.Abs(c1.real)
		c1.imaginary = math.Abs(c1.imaginary)
	}
	result := sum(power(c1, orderOfFractalExpression), c)
	c1.setEqualTo(result)
}

//func zRecursive(iteration int, c *complex) *complex {
//	if iteration == 0 {
//		return newComplex(0, 0)
//	}
//	// zRecursive[n-1] ** 2 + c
//	new := zRecursive(iteration-1, c)
//	return sum(multiply(new, new), c)
//}

func recurrentIterationsBeforeBlowingUp(z, c *complex) int {
	i := 0
	for ; i < maxSetCheckIterations; i++ {
		z.iterateAsZ(c)
		if z.squareMagnitude() > 4 {
			return i
		}
	}
	return -1 // not blowing up
}

func isPartOfMandelbrotForPrecision(c *complex, precision int) bool {
	t := maxSetCheckIterations
	maxSetCheckIterations = precision + 2
	iters := recurrentIterationsBeforeBlowingUp(newComplex(0, 0), c)
	// fmt.Printf("Current iters: %d", iters)
	maxSetCheckIterations = t
	return iters >= precision
}

func getMandelbrotIterations(c *complex) int {
	return recurrentIterationsBeforeBlowingUp(newComplex(0, 0), c)
}

func isPartOfJulia(candidate, parameter *complex) bool {
	return recurrentIterationsBeforeBlowingUp(candidate, parameter) < maxSetCheckIterations
}

func getJuliaIterations(candidate, parameter *complex) int {
	return recurrentIterationsBeforeBlowingUp(candidate, parameter)
}

func generateNewJuliaParameter() {
	parameter := randomComplex(-2, 2, -2, 2)
	juliaParameter.setEqualTo(parameter)
	for !isPartOfMandelbrotForPrecision(parameter, rand.Intn(500)+5) {
		parameter = randomComplex(-2, 2, -2, 2)
		juliaParameter.setEqualTo(parameter)
	}
}

func getIterationsForCurrentFractalType(c, p *complex) int {
	switch currentFractal {
	case MANDELBROT, TRICORN, BURNING_SHIP:
		return getMandelbrotIterations(c)
	case JULIA_MB, JULIA_TRICORN, JULIA_BS:
		return getJuliaIterations(c, p)
	}
	panic("Unknown fractal type")
}

func getCurrFractalNameString() string {
	name := ""
	switch currentFractal {
	case MANDELBROT:
		name = "Mandelbrot"
	case JULIA_MB:
		name = "Julia for mandelbrot"
	case TRICORN:
		name = "Tricorn"
	case JULIA_TRICORN:
		name = "Julia for tricorn"
	case BURNING_SHIP:
		name = "Burning ship"
	case JULIA_BS:
		name = "Julia for burning ship"
	default:
		panic("Unknown fractal number")
	}
	return name
}
