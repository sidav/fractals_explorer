package main

var maxSetCheckIterations = 20
var orderOfFractalExpression = 2

func (c1 *complex) iterateAsZ(c *complex) {
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
