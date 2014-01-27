package automatehw

import (
	//"fmt"
	"math"
	"math/rand"
	"strconv"
)

//Generate random equation initialized by the fn of the student
func RandomHomogenousEquations(fn *string, variables []string, numberOfEquations, maxnumber int) []string {
	//initialize the random generator
	num, _ := strconv.ParseInt(*fn, 10, 0)
	source := rand.NewSource(num)
	random := rand.New(source)
	//declare the variable we need
	var homogenousEquations []string
	var homogenousEquation string
	var coefficient int

	for count := 0; count < numberOfEquations; count++ {
		homogenousEquation = ""
		for i, value := range variables {
			//geneate a coeficient with value from 1 to maxnumber, multiplied by -1 powered by random number
			coefficient = (random.Intn(maxnumber-1) + 1) * int(math.Pow(-1, float64(random.Int())))
			//if this is the first coefficient or the coefficient is negative we do not add a sign
			//otherways we add a + sign infront of it.
			if (i == 0) || (coefficient < 0) {
				homogenousEquation = homogenousEquation + strconv.Itoa(coefficient) + "*" + string(value)
			} else {
				homogenousEquation = homogenousEquation + "+" + strconv.Itoa(coefficient) + "*" + string(value)
			}
		}
		//finally we add the equal sign and append the equation to the system
		homogenousEquation = homogenousEquation + "=0"
		homogenousEquations = append(homogenousEquations, homogenousEquation)
	}
	return homogenousEquations
}

// Generate homogenous equations with coeficients the figures of the FN
func GenerateHomogenousEquationsFromFN(fn *string, variables []string, numberOfEquations, maxnumber int) []string {
	//initialize the random generator
	num, _ := strconv.ParseInt(*fn, 10, 0)
	source := rand.NewSource(num)
	random := rand.New(source)

	var homogenousEquations, permutation []string
	var homogenousEquation string
	var coefficient, rnum int

	for count := 0; count < numberOfEquations; count++ {
		homogenousEquation = ""
		permutation = RandomPermutationOfFN(fn, count*13)
		for i, value := range variables {
			rnum, _ = strconv.Atoi(permutation[i])
			for rnum == 0 {
				rnum = random.Intn(maxnumber)
			}
			coefficient = rnum * int(math.Pow(-1, float64(random.Int())))
			if i == 0 {
				homogenousEquation = homogenousEquation + strconv.Itoa(coefficient) + "*" + string(value)
			} else {
				if coefficient < 0 {
					homogenousEquation = homogenousEquation + strconv.Itoa(coefficient) + "*" + string(value)
				} else {
					homogenousEquation = homogenousEquation + "+" + strconv.Itoa(coefficient) + "*" + string(value)
				}
			}
		}
		homogenousEquation = homogenousEquation + "=0"
		homogenousEquations = append(homogenousEquations, homogenousEquation)
	}
	return homogenousEquations
}

//Generate random permutation of the figures in the FN
func RandomPermutationOfFN(fn *string, n int) []string {
	//initialize the random generator
	num, _ := strconv.ParseInt(*fn, 10, 0)
	source := rand.NewSource(num * int64(n))
	random := rand.New(source)
	//we generate a random permutation of the numbers from 0 to len(fn),
	//which we use as indexes for the fn string, geting a random permutation of the fn.
	intPerm := random.Perm(len(*fn))
	var permutation []string
	for _, value := range intPerm {
		permutation = append(permutation, string((*fn)[value]))
	}
	return permutation
}

//Generate array of int arrays - each representing a vector. The function receives
//as input parameters dim - the dimenssion of the vectors, n - the number of vectors to generate,
//max the maximum value of each of the coordinates of the vectors and positive - a bool flag, representing
//whether the values of the coordinates of the vector should be positive or not.
func GenerateVectors(fn *string, dim, n, max int, positive bool) [][]int {
	var vectors [][]int
	var vector []int

	//initialize the random generator
	num, _ := strconv.ParseInt(*fn, 10, 0)
	source := rand.NewSource(num)
	random := rand.New(source)

	//Here begins the case, when the coordinates are positive.
	if positive == true {
		for i := 0; i < n; i++ {
			vector = []int{}
			if i > 0 {
				//For each generated vector after the first one
				//we check, whether it is linear independent of the previous vectors.
				for IsLinearIndependent(vectors[0], vector) {
					for t := 0; t < dim; t++ {
						value := random.Intn(max-1) + 1
						vector = append(vector, value)
					}
				}
			} else {
				//Here we generate the first vector
				for t := 0; t < dim; t++ {
					value := random.Intn(max-1) + 1
					vector = append(vector, value)
				}
			}
			//finally we append the generated vector
			vectors = append(vectors, vector)
		}
	} else {
		//If the vector could have negative coordinates, we proceed the same way,
		//but this time we multiply the vector coordinates with (-1)^random_number
		for i := 0; i < n; i++ {
			vector = []int{0, 0, 0, 0}
			if i > 0 {
				//For each vector after the first one we check whether it is linear independent to the previous
				for !IsLinearIndependent(vectors[0], vector) {
					vector = []int{}
					for t := 0; t < dim; t++ {
						value := (random.Intn(max-1) + 1) * int(math.Pow(-1, float64(random.Intn(13))))
						vector = append(vector, value)
					}
				}
			} else {
				vector = []int{}
				//Here we generate the first vector
				for t := 0; t < dim; t++ {
					value := random.Intn(max-1) + 1
					vector = append(vector, value)
				}
			}
			// Here we append the generated vector
			vectors = append(vectors, vector)
		}
	}
	return vectors
}
